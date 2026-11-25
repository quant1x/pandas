const std = @import("std");
const support = @import("support.zig");

/// REF 引用前 N 个周期的值；越界为 NaN。
pub fn refValue(allocator: std.mem.Allocator, x: []const f64, n: i64) ![]f64 {
    const out = try allocator.alloc(f64, x.len);
    if (n == 0) {
        @memcpy(out, x);
        return out;
    }
    @memset(out, support.nan());
    const distance: i128 = if (n < 0) -@as(i128, n) else @as(i128, n);
    const width: usize = @intCast(@min(distance, @as(i128, @intCast(x.len))));
    if (n > 0) {
        @memcpy(out[width..], x[0 .. x.len - width]);
    } else {
        @memcpy(out[0 .. x.len - width], x[width..]);
    }
    return out;
}
