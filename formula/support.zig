//! 共用辅助（Go 侧由 pandas 的 Series.Rolling 承担，不在 formula/ 目录）。

const std = @import("std");

pub fn nan() f64 {
    return std.math.nan(f64);
}

/// 逐窗口应用：完整窗口才计算，窗口不足为 NaN。
pub fn rolling(
    comptime apply: fn ([]const f64) f64,
    allocator: std.mem.Allocator,
    x: []const f64,
    window: i64,
) ![]f64 {
    const out = try allocator.alloc(f64, x.len);
    @memset(out, nan());
    if (window <= 0) return out;
    const size: usize = @intCast(window);
    if (size > x.len) return out;
    var index: usize = size - 1;
    while (index < x.len) : (index += 1) {
        out[index] = apply(x[index + 1 - size .. index + 1]);
    }
    return out;
}

pub fn totalOf(block: []const f64) f64 {
    var total: f64 = 0;
    for (block) |value| total += value;
    return total;
}

/// EWM `Adjust=false` 递归：首值 NaN 视为 0，NaN 位置沿用前值。
pub fn ewm(allocator: std.mem.Allocator, x: []const f64, alpha: f64) ![]f64 {
    const out = try allocator.alloc(f64, x.len);
    if (x.len == 0) return out;
    var previous = x[0];
    if (std.math.isNan(previous)) previous = 0;
    out[0] = previous;
    const beta = 1 - alpha;
    var index: usize = 1;
    while (index < x.len) : (index += 1) {
        const value = x[index];
        if (std.math.isNan(value)) {
            out[index] = previous;
            continue;
        }
        previous = beta * previous + alpha * value;
        if (std.math.isNan(previous)) previous = out[index - 1];
        out[index] = previous;
    }
    return out;
}
