const std = @import("std");
const support = @import("support.zig");

/// LLV 最近 N 个周期的最小值。
pub fn llv(allocator: std.mem.Allocator, x: []const f64, window: i64) ![]f64 {
    return support.rolling(minOf, allocator, x, window);
}

fn minOf(block: []const f64) f64 {
    var best = block[0];
    for (block[1..]) |value| {
        if (value < best) best = value;
    }
    return best;
}
