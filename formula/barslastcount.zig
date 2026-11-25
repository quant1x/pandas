const std = @import("std");
const support = @import("support.zig");

/// BARSLASTCOUNT 连续成立的周期数；遇到不成立归零。
pub fn barslastcount(allocator: std.mem.Allocator, condition: []const f64) ![]f64 {
    const out = try allocator.alloc(f64, condition.len);
    var streak: f64 = 0;
    for (condition, 0..) |value, index| {
        streak = if (value != 0) streak + 1 else 0;
        out[index] = streak;
    }
    return out;
}
