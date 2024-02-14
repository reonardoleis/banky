const runHttpServer = @import("http_server.zig").runHttpServer;
pub fn main() !void {
    try runHttpServer();
}
