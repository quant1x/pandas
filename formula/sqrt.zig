const std = @import("std");
const support = @import("support.zig");

/// SQRT 逐元素平方根。
pub fn sqrt(allocator: std.mem.Allocator, x: []const f64) ![]f64 {
    const out = try allocator.alloc(f64, x.len);
    for (x, 0..) |value, index| {
        out[index] = @sqrt(value);
    }
    return out;
}
