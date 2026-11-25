const std = @import("std");
const support = @import("support.zig");

/// ABS 逐元素绝对值。
pub fn abs(allocator: std.mem.Allocator, x: []const f64) ![]f64 {
    const out = try allocator.alloc(f64, x.len);
    for (x, 0..) |value, index| {
        out[index] = @abs(value);
    }
    return out;
}
