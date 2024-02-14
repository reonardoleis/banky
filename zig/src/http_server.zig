const std = @import("std");
const handlers = @import("handlers.zig");
pub fn runHttpServer() !void {
    const server_addr = "127.0.0.1";
    const server_port = 8080;

    var gpa = std.heap.GeneralPurposeAllocator(.{}){};

    defer std.debug.assert(gpa.deinit() == .ok);

    var allocator = gpa.allocator();

    var server = std.http.Server.init(allocator, .{ .reuse_address = true });
    defer server.deinit();

    std.log.info("server running at ${s}:${d}", .{ server_addr, server_port });

    const addr = std.net.Address.parseIp(server_addr, server_port) catch unreachable;

    try server.listen(addr);

    outer: while (true) {
        var response = try server.accept(.{
            .allocator = allocator,
        });

        defer response.deinit();

        while (response.reset() != .closing) {
            response.wait() catch |err| switch (err) {
                error.HttpHeadersInvalid => continue :outer,
                error.EndOfStream => continue,
                else => return err,
            };
        }

        try handleRequest(&response, allocator);
    }
}

fn handleRequest(response: *std.http.Server.Response, allocator: std.mem.Allocator) !void {
    const log = std.log;
    log.info("{s} {s} {s}", .{ @tagName(response.request.method), @tagName(response.request.version), response.request.target });

    const body = try response.reader().readAllAlloc(allocator, 8192);
    defer allocator.free(body);

    if (response.request.headers.contains("connection")) {
        try response.headers.append("connection", "keep-alive");
    }

    try switch (response.request.method) {
        std.http.Method.GET => handlers.handleRequestStatement(response, allocator),
        else => handlers.handleDispatchCreation(response, allocator),
    };
}
