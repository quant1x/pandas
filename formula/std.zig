const std = @import("std");
const support = @import("support.zig");

/// STD 最近 N 个周期的总体标准差（符号名 `stddev` 避开 `std`）。
pub fn stddev(allocator: std.mem.Allocator, x: []const f64, window: i64) ![]f64 {
    return support.rolling(deviationOf, allocator, x, window);
}

fn deviationOf(block: []const f64) f64 {
    const size: f64 = @floatFromInt(block.len);
    const mean = support.totalOf(block) / size;
    var squares: f64 = 0;
    for (block) |value| {
        const delta = value - mean;
        squares += delta * delta;
    }
    return @sqrt(squares / size);
}
