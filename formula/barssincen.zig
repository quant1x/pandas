const std = @import("std");
const support = @import("support.zig");

/// BARSSINCEN N 周期内第一次成立到当前的周期数；窗口不足为 0。
pub fn barssincen(allocator: std.mem.Allocator, condition: []const f64, window: i64) ![]f64 {
    const out = try allocator.alloc(f64, condition.len);
    @memset(out, 0);
    if (window <= 0) return out;
    const size: usize = @intCast(window);
    if (size > condition.len) return out;
    var index: usize = size - 1;
    while (index < condition.len) : (index += 1) {
        const begin = index + 1 - size;
        for (condition[begin .. index + 1], begin..) |value, position| {
            if (value != 0) {
                out[index] = @floatFromInt(index - position);
                break;
            }
        }
    }
    return out;
}
