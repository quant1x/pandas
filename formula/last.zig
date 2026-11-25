const std = @import("std");
const support = @import("support.zig");

/// LAST 持续存在：A+1 窗口内排除最近 B 根后全部非零；窗口不足为 NaN。
pub fn last(allocator: std.mem.Allocator, condition: []const f64, a: i64, b: i64) ![]f64 {
    const out = try allocator.alloc(f64, condition.len);
    @memset(out, support.nan());
    if (a < 0) return out;
    const size: usize = @intCast(a + 1);
    if (size == 0 or size > condition.len) return out;
    const skip: usize = if (b < 0) 0 else @intCast(b);
    var index: usize = size - 1;
    while (index < condition.len) : (index += 1) {
        if (skip >= size) {
            out[index] = 1;
            continue;
        }
        const begin = index + 1 - size;
        const stop = index + 1 - skip;
        var satisfied = true;
        for (condition[begin..stop]) |value| {
            if (value == 0) {
                satisfied = false;
                break;
            }
        }
        out[index] = if (satisfied) 1 else 0;
    }
    return out;
}
