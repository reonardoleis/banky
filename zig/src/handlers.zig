const std = @import("std");

pub fn handleDispatchCreation(response: *std.http.Server.Response, allocator: std.mem.Allocator) !void {
    const body = try response.reader().readAllAlloc(allocator, 8192);
    defer allocator.free(body);

    response.transfer_encoding = .{ .content_length = 8 };

    try response.headers.append("content-type", "text/plain");

    try response.do();
    if (response.request.method != .HEAD) {
        try response.writeAll("creation");
        try response.finish();
    }
}

pub fn handleRequestStatement(response: *std.http.Server.Response, allocator: std.mem.Allocator) !void {
    const body = try response.reader().readAllAlloc(allocator, 8192);
    defer allocator.free(body);

    response.transfer_encoding = .{ .content_length = 9 };

    try response.headers.append("content-type", "text/plain");

    try response.do();
    if (response.request.method != .HEAD) {
        try response.writeAll("statement");
        try response.finish();
    }
}
