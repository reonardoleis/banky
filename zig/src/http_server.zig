const std = @import("std");
const zap = @import("zap");
const handlers = @import("handlers.zig");
fn on_request(r: zap.Request) void {
    const method = r.method.?;

    if (method[0] == 'G') {
        handlers.statement(r);
    } else {
        handlers.create(r);
    }
}

pub fn start() !void {
    var listener = zap.HttpListener.init(.{
        .port = 8080,
        .on_request = on_request,
        .log = true,
    });

    try listener.listen();

    std.debug.print("listening on localhost:8080\n", .{});

    zap.start(.{
        .threads = 4,
        .workers = 4,
    });
}
