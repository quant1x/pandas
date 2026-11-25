const std = @import("std");
const support = @import("support.zig");

/// SMA 中国式平滑：alpha = M/N（M=0 视为 1）。
pub fn sma(allocator: std.mem.Allocator, x: []const f64, n: i64, m: i64) ![]f64 {
    const weight: f64 = @floatFromInt(if (m == 0) @as(i64, 1) else m);
    const period: f64 = @floatFromInt(@max(n, 1));
    return support.ewm(allocator, x, weight / period);
}
