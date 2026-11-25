const std = @import("std");
const support = @import("support.zig");

/// LLVBARS 最近 N 周期内最低点到当前的周期数；并列取最近的一个。
pub fn llvbars(allocator: std.mem.Allocator, x: []const f64, window: i64) ![]f64 {
    return support.rolling(barsSinceMin, allocator, x, window);
}

fn barsSinceMin(block: []const f64) f64 {
    var position: usize = 0;
    var best = block[0];
    for (block, 0..) |value, index| {
        if (value <= best) {
            best = value;
            position = index;
        }
    }
    return @floatFromInt(block.len - 1 - position);
}
