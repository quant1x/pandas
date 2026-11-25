const std = @import("std");
const support = @import("support.zig");

/// CROSS 金叉：前一根 S1 < S2 且当前 S1 > S2。
pub fn cross(allocator: std.mem.Allocator, s1: []const f64, s2: []const f64) ![]bool {
    const out = try allocator.alloc(bool, s1.len);
    @memset(out, false);
    const length = @min(s1.len, s2.len);
    var index: usize = 1;
    while (index < length) : (index += 1) {
        out[index] = s1[index - 1] < s2[index - 1] and s1[index] > s2[index];
    }
    return out;
}
