const std = @import("std");
const support = @import("support.zig");

/// SUM 最近 N 个周期的和。
pub fn sum(allocator: std.mem.Allocator, x: []const f64, window: i64) ![]f64 {
    return support.rolling(support.totalOf, allocator, x, window);
}
