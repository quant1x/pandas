const std = @import("std");
const support = @import("support.zig");

/// DMA 动态移动平均：alpha = 1/A，A 可常量或逐位置；非有限时按 1。
pub fn dma(allocator: std.mem.Allocator, x: []const f64, a: []const f64) ![]f64 {
    const out = try allocator.alloc(f64, x.len);
    if (x.len == 0) return out;
    out[0] = x[0];
    var index: usize = 1;
    while (index < x.len) : (index += 1) {
        const period = if (a.len == 1) a[0] else a[index];
        var alpha = 1.0 / period;
        if (!std.math.isFinite(alpha)) alpha = 1.0;
        out[index] = alpha * x[index] + (1 - alpha) * out[index - 1];
    }
    return out;
}
