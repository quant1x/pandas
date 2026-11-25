const std = @import("std");
const support = @import("support.zig");

/// MIN 两条序列逐元素取小；要求等长。
pub fn minSeries(allocator: std.mem.Allocator, x: []const f64, y: []const f64) ![]f64 {
    if (x.len != y.len) return error.LengthMismatch;
    const out = try allocator.alloc(f64, x.len);
    for (x, y, 0..) |left, right, index| {
        out[index] = @min(left, right);
    }
    return out;
}
