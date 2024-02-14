const zap = @import("zap");
const std = @import("std");
const models = @import("models.zig");
const dto = @import("dto.zig");

pub fn statement(r: zap.Request) void {
    var it = std.mem.split(u8, r.path.?, "/");
    _ = it.next();
    _ = it.next();
    // var user_id = it.next();

    var transactions: []models.Transaction = undefined;
    var response = dto.toStatementResponse(200, 200, transactions);
    var buf: [256]u8 = undefined;
    var json_to_send: []const u8 = undefined;

    if (zap.stringifyBuf(&buf, response, .{})) |json| {
        json_to_send = json;
    } else {
        json_to_send = "null";
    }

    r.setContentType(.JSON) catch return;
    r.sendBody(json_to_send) catch return;
}

pub fn create(r: zap.Request) void {
    var response = dto.toCreateResponse(200, 200);
    var buf: [256]u8 = undefined;
    var json_to_send: []const u8 = undefined;

    if (zap.stringifyBuf(&buf, response, .{})) |json| {
        json_to_send = json;
    } else {
        json_to_send = "null";
    }

    r.setContentType(.JSON) catch return;
    r.sendBody(json_to_send) catch return;
}
