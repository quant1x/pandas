const std = @import("std");
const support = @import("support.zig");

/// BARSLAST 上一次条件成立到当前的周期数；从未成立为 -1。
pub fn barslast(allocator: std.mem.Allocator, condition: []const f64) ![]f64 {
    const out = try allocator.alloc(f64, condition.len);
    var last_true: ?usize = null;
    for (condition, 0..) |value, index| {
        if (value != 0) {
            last_true = index;
            out[index] = 0;
        } else if (last_true) |position| {
            out[index] = @floatFromInt(index - position);
        } else {
            out[index] = -1;
        }
    }
    return out;
}
