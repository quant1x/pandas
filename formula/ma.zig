const std = @import("std");
const support = @import("support.zig");

/// MA 简单移动平均：完整窗口才计算，窗口不足为 NaN。
pub fn ma(allocator: std.mem.Allocator, x: []const f64, window: i64) ![]f64 {
    const out = try allocator.alloc(f64, x.len);
    @memset(out, support.nan());
    if (window <= 0) return out;
    const size: usize = @intCast(window);
    if (size > x.len) return out;
    var index: usize = size - 1;
    while (index < x.len) : (index += 1) {
        const start = index + 1 - size;
        var total: f64 = 0;
        for (x[start .. index + 1]) |value| {
            total += value;
        }
        out[index] = total / @as(f64, @floatFromInt(size));
    }
    return out;
}
