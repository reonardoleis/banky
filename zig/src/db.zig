const pg = @import("pg");
const std = @import("std");

pub fn init() void {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    const allocator = gpa.allocator();

    var pool = try pg.Pool.init(allocator, .{
        .size = 10,
        .connect = .{
            .port = 5432,
            .host = "localhost",
        },
        .auth = .{
            .username = "postgres",
            .database = "banky",
            .password = "postgres",
            .timeout = 10_000,
        },
    });

    std.debug.log("pool type {any}\n", .{@TypeOf(pool)});

    defer pool.deinit();
    defer gpa.deinit();
}
