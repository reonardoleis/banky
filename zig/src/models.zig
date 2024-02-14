pub const Account = struct {
    id: i32 = 0,
    limit: i64 = 0,
    initial_balance: i64 = 0,
};

pub const Transaction = struct {
    amount: i32 = 0,
    type: u8 = 0,
    description: ?[]const u8 = null,
};
