const std = @import("std");
const support = @import("support.zig");

/// WMA 加权移动平均：窗口内按 1..N 递增加权，分母 N(N+1)/2。
pub fn wma(allocator: std.mem.Allocator, x: []const f64, window: i64) ![]f64 {
    return support.rolling(weightedMeanOf, allocator, x, window);
}

fn weightedMeanOf(block: []const f64) f64 {
    var total: f64 = 0;
    var weights: f64 = 0;
    for (block, 0..) |value, index| {
        const weight: f64 = @floatFromInt(index + 1);
        total += weight * value;
        weights += weight;
    }
    return total / weights;
}
