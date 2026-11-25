const std = @import("std");
const support = @import("support.zig");

/// AVEDEV 平均绝对偏差：窗口内「与均值之差的绝对值」的平均。
pub fn avedev(allocator: std.mem.Allocator, x: []const f64, window: i64) ![]f64 {
    return support.rolling(absoluteDeviationOf, allocator, x, window);
}

fn absoluteDeviationOf(block: []const f64) f64 {
    const size: f64 = @floatFromInt(block.len);
    const mean = support.totalOf(block) / size;
    var total: f64 = 0;
    for (block) |value| total += @abs(value - mean);
    return total / size;
}
