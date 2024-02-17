const http_server = @import("http_server.zig");
const db = @import("db.zig");
pub fn main() !void {
    try db.init();
    try http_server.start();
}
