const std = @import("std");
const support = @import("support.zig");

/// IF 条件选择：条件非零取 A，否则取 B；三条序列等长。
pub fn ifSelect(
    allocator: std.mem.Allocator,
    condition: []const f64,
    a: []const f64,
    b: []const f64,
) ![]f64 {
    if (condition.len != a.len or condition.len != b.len) return error.LengthMismatch;
    const out = try allocator.alloc(f64, condition.len);
    for (condition, a, b, 0..) |flag, left, right, index| {
        out[index] = if (flag != 0) left else right;
    }
    return out;
}

/// IFN 反向条件选择：条件非零取 B，否则取 A。
pub fn ifn(
    allocator: std.mem.Allocator,
    condition: []const f64,
    a: []const f64,
    b: []const f64,
) ![]f64 {
    return ifSelect(allocator, condition, b, a);
}
