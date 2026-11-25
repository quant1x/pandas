const std = @import("std");
const support = @import("support.zig");

/// HHV 最近 N 个周期的最大值。
pub fn hhv(allocator: std.mem.Allocator, x: []const f64, window: i64) ![]f64 {
    return support.rolling(maxOf, allocator, x, window);
}

fn maxOf(block: []const f64) f64 {
    var best = block[0];
    for (block[1..]) |value| {
        if (value > best) best = value;
    }
    return best;
}
