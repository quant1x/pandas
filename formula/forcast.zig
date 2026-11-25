const std = @import("std");
const support = @import("support.zig");

const slope = @import("slope.zig");

/// FORCAST 窗口内线性回归在最后一根上的预测值。
pub fn forcast(allocator: std.mem.Allocator, x: []const f64, window: i64) ![]f64 {
    return support.rolling(forecastOf, allocator, x, window);
}

fn forecastOf(block: []const f64) f64 {
    const line = slope.linearRegression(block);
    if (std.math.isNan(line.slope)) return support.nan();
    return line.intercept + line.slope * @as(f64, @floatFromInt(block.len - 1));
}
