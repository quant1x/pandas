const std = @import("std");
const support = @import("support.zig");

const Line = struct { slope: f64, intercept: f64 };

/// SLOPE 窗口内线性回归斜率；不足两点为 NaN。
pub fn slope(allocator: std.mem.Allocator, x: []const f64, window: i64) ![]f64 {
    return support.rolling(slopeOf, allocator, x, window);
}

pub fn linearRegression(block: []const f64) Line {
    if (block.len < 2) return .{ .slope = support.nan(), .intercept = support.nan() };
    const size: f64 = @floatFromInt(block.len);
    const mean_t = (size - 1) / 2;
    var mean_y: f64 = 0;
    for (block) |value| mean_y += value;
    mean_y /= size;
    var numerator: f64 = 0;
    var denominator: f64 = 0;
    for (block, 0..) |value, index| {
        const t = @as(f64, @floatFromInt(index)) - mean_t;
        numerator += t * (value - mean_y);
        denominator += t * t;
    }
    const fitted = numerator / denominator;
    return .{ .slope = fitted, .intercept = mean_y - fitted * mean_t };
}

fn slopeOf(block: []const f64) f64 {
    return linearRegression(block).slope;
}
