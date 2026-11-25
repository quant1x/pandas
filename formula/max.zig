const std = @import("std");
const support = @import("support.zig");

/// MAX 两条序列逐元素取大；要求等长。
pub fn maxSeries(allocator: std.mem.Allocator, x: []const f64, y: []const f64) ![]f64 {
    if (x.len != y.len) return error.LengthMismatch;
    const out = try allocator.alloc(f64, x.len);
    for (x, y, 0..) |left, right, index| {
        out[index] = @max(left, right);
    }
    return out;
}
