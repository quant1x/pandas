const std = @import("std");
const support = @import("support.zig");

/// COUNT 最近 N 个周期内非零条件的个数；窗口不足计 0。
pub fn count(allocator: std.mem.Allocator, condition: []const f64, window: i64) ![]f64 {
    const out = try allocator.alloc(f64, condition.len);
    @memset(out, 0);
    if (window <= 0) return out;
    const size: usize = @intCast(window);
    if (size > condition.len) return out;
    var index: usize = size - 1;
    while (index < condition.len) : (index += 1) {
        var total: usize = 0;
        for (condition[index + 1 - size .. index + 1]) |value| {
            if (value != 0) total += 1;
        }
        out[index] = @floatFromInt(total);
    }
    return out;
}
