const std = @import("std");
const support = @import("support.zig");

/// EMA 指数移动平均：alpha = 2/(N+1)，首值等于源序列首值。
pub fn ema(allocator: std.mem.Allocator, x: []const f64, span: i64) ![]f64 {
    const period: f64 = @floatFromInt(@max(span, 1));
    return support.ewm(allocator, x, 2.0 / (period + 1.0));
}
