const std = @import("std");
const support = @import("support.zig");

/// CONST 取序列最后值为常量序列；Go 先转 float32，故按 f32 精度落值。
pub fn constSeries(allocator: std.mem.Allocator, x: []const f64) ![]f64 {
    const out = try allocator.alloc(f64, x.len);
    if (x.len == 0) return out;
    const narrowed: f32 = @floatCast(x[x.len - 1]);
    @memset(out, @as(f64, narrowed));
    return out;
}
