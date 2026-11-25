const std = @import("std");
const support = @import("support.zig");

/// BARSLASTS 倒数第 N 次成立到当前的周期数；越界为 NaN。
pub fn barslasts(allocator: std.mem.Allocator, condition: []const f64, n: i64) ![]f64 {
    var current = try barslastZero(allocator, condition);
    var iteration: i64 = 1;
    while (iteration < n) : (iteration += 1) {
        const following = try allocator.alloc(f64, current.len);
        @memset(following, support.nan());
        for (current, 0..) |value, index| {
            if (std.math.isNan(value)) continue;
            const step = value + 1;
            const source = @as(f64, @floatFromInt(index)) - step;
            if (source < 0) continue;
            const source_index: usize = @intFromFloat(source);
            if (source_index >= current.len) continue;
            if (std.math.isNan(current[source_index])) continue;
            following[index] = current[source_index] + step;
        }
        allocator.free(current);
        current = following;
    }
    return current;
}

fn barslastZero(allocator: std.mem.Allocator, condition: []const f64) ![]f64 {
    const out = try allocator.alloc(f64, condition.len);
    var previous: f64 = 0;
    for (condition, 0..) |value, index| {
        previous = if (value != 0) 0 else previous + 1;
        out[index] = previous;
    }
    return out;
}
