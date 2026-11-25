const std = @import("std");
const support = @import("support.zig");

/// FILTER 条件成立后 N 位清零；扫描已处理的序列。
pub fn filter(allocator: std.mem.Allocator, condition: []const f64, window: i64) ![]f64 {
    const out = try allocator.alloc(f64, condition.len);
    @memcpy(out, condition);
    if (window <= 0) return out;
    const width: usize = @intCast(window);
    var index: usize = 0;
    while (index < out.len) : (index += 1) {
        if (out[index] == 0) continue;
        const end = @min(index + 1 + width, out.len);
        @memset(out[index + 1 .. end], 0);
    }
    return out;
}
