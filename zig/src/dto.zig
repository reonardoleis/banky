const models = @import("models.zig");

pub const CreateRequest = struct {
    valor: i32 = 0,
    tipo: u8 = 0,
    descricao: ?[]const u8 = null,
};

pub const CreateResponse = struct {
    limite: i32 = 0,
    saldo: i32 = 0,
};

pub fn toCreateResponse(limit: i32, balance: i32) CreateResponse {
    return .{
        .limite = limit,
        .saldo = balance,
    };
}

// pub const StatementRequest (id from path param)

pub const StatementBalance = struct {
    total: i32 = 0,
    data_extrato: ?[]const u8 = null,
    limite: i32 = 0,
};

pub const StatementTransaction = struct {
    valor: i32 = 0,
    tipo: u8 = 0,
    descricao: ?[]const u8 = null,
    realizada_em: ?[]const u8 = null,
};

pub const StatementResponse = struct {
    saldo: ?StatementBalance = null,
    ultimas_transacoes: ?[]StatementTransaction = null,
};

pub fn toStatementResponse(total: i32, limit: i32, transactions: []models.Transaction) StatementResponse {
    _ = transactions;
    return .{
        .saldo = .{
            .total = total,
            .data_extrato = "2024",
            .limite = limit,
        },
    };
}
