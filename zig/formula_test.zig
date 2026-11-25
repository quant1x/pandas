//! 公式族 Zig 实现的单元用例与 CSV fixture 对拍。
//!
//! 在 `pandas/` 目录下执行 `zig test zig/formula_test.zig`；fixture 位于
//! `tests/formula/cases.csv`，与 Go / C++ / Rust 共用同一份期望值。

const std = @import("std");
const formula = @import("formula.zig");

fn nan() f64 {
    return std.math.nan(f64);
}

/// 读取 fixture：先按项目根目录解析，再按 `zig/` 目录解析。
fn readFixture(allocator: std.mem.Allocator) ![]u8 {
    const dir = std.Io.Dir.cwd();
    const limit = std.Io.Limit.limited(1 << 20);
    if (dir.readFileAlloc(std.testing.io, "tests/formula/cases.csv", allocator, limit)) |text| {
        return text;
    } else |_| {}
    return dir.readFileAlloc(std.testing.io, "../tests/formula/cases.csv", allocator, limit);
}

fn parseList(allocator: std.mem.Allocator, text: []const u8) ![]f64 {
    var list: std.ArrayList(f64) = .empty;
    errdefer list.deinit(allocator);
    var parts = std.mem.tokenizeScalar(u8, text, ';');
    while (parts.next()) |part| {
        const item = std.mem.trim(u8, part, " \t\r");
        const value = if (std.mem.eql(u8, item, "nan"))
            nan()
        else
            try std.fmt.parseFloat(f64, item);
        try list.append(allocator, value);
    }
    return list.toOwnedSlice(allocator);
}

/// 布尔序列在 fixture 里用 0/1 表示。
fn boolsAsFloat(allocator: std.mem.Allocator, flags: []const bool) ![]f64 {
    const out = try allocator.alloc(f64, flags.len);
    for (flags, 0..) |flag, index| {
        out[index] = if (flag) 1.0 else 0.0;
    }
    return out;
}

/// 0/1 序列转布尔序列；非零即真。
fn floatsAsBools(allocator: std.mem.Allocator, values: []const f64) ![]bool {
    const out = try allocator.alloc(bool, values.len);
    for (values, 0..) |value, index| {
        out[index] = value != 0;
    }
    return out;
}

/// NOT 的 fixture 包装：输入 0/1 序列。
fn compareNot(allocator: std.mem.Allocator, x: []const f64) ![]f64 {
    const values = try floatsAsBools(allocator, x);
    defer allocator.free(values);
    const flags = try formula.logicNot(allocator, values);
    defer allocator.free(flags);
    return boolsAsFloat(allocator, flags);
}

/// IF / IFN 的 fixture 包装：B 支用常量序列填充。
fn ifBranch(
    allocator: std.mem.Allocator,
    condition: []const f64,
    values: []const f64,
    fallback: f64,
    invert: bool,
) ![]f64 {
    const filler = try allocator.alloc(f64, condition.len);
    defer allocator.free(filler);
    @memset(filler, fallback);
    return if (invert)
        formula.ifn(allocator, condition, values, filler)
    else
        formula.ifSelect(allocator, condition, values, filler);
}

/// CROSS 返回布尔序列，fixture 用 0/1 表示。
fn crossAsFloat(allocator: std.mem.Allocator, x: []const f64, y: []const f64) ![]f64 {
    const flags = try formula.cross(allocator, x, y);
    defer allocator.free(flags);
    return boolsAsFloat(allocator, flags);
}

/// 比较族返回布尔序列，fixture 用 0/1 表示。
fn compareAsFloat(
    comptime op: fn (std.mem.Allocator, []const f64, []const f64) anyerror![]bool,
    allocator: std.mem.Allocator,
    x: []const f64,
    y: []const f64,
) ![]f64 {
    const flags = try op(allocator, x, y);
    defer allocator.free(flags);
    return boolsAsFloat(allocator, flags);
}

/// 逻辑族返回布尔序列，fixture 用 0/1 表示。
fn logicBinaryAsFloat(
    comptime op: fn (std.mem.Allocator, []const bool, []const bool) anyerror![]bool,
    allocator: std.mem.Allocator,
    x: []const f64,
    y: []const f64,
) ![]f64 {
    const left = try floatsAsBools(allocator, x);
    defer allocator.free(left);
    const right = try floatsAsBools(allocator, y);
    defer allocator.free(right);
    const flags = try op(allocator, left, right);
    defer allocator.free(flags);
    return boolsAsFloat(allocator, flags);
}

fn parsePeriod(text: []const u8) !i64 {
    if (text.len == 0) return 0;
    return std.fmt.parseInt(i64, std.mem.trim(u8, text, " \t\r"), 10);
}

/// 逐值比对：NaN 对 NaN 视为相等，其余按相对 1e-12（下限 1）容差。
fn expectListEqual(name: []const u8, got: []const f64, want: []const f64) !void {
    if (got.len != want.len) {
        std.debug.print("{s}: 长度 {d} != {d}\n", .{ name, got.len, want.len });
        return error.LengthMismatch;
    }
    for (got, want, 0..) |actual, expected, index| {
        if (std.math.isNan(actual) and std.math.isNan(expected)) continue;
        const tolerance = 1e-12 * @max(@abs(expected), 1.0);
        if (@abs(actual - expected) <= tolerance) continue;
        std.debug.print("{s}: 第 {d} 位 {d} != {d}\n", .{ name, index, actual, expected });
        return error.ValueMismatch;
    }
}

test "abs 逐元素绝对值" {
    const allocator = std.testing.allocator;
    const x = [_]f64{ 1.0, -2.0, 3.0, -0.0, nan() };
    const got = try formula.abs(allocator, &x);
    defer allocator.free(got);
    try expectListEqual("abs", got, &[_]f64{ 1.0, 2.0, 3.0, 0.0, nan() });
}

test "ref 越界为 NaN" {
    const allocator = std.testing.allocator;
    const x = [_]f64{ 1.0, 2.0, 3.0, 4.0, 5.0 };
    const got = try formula.refValue(allocator, &x, 2);
    defer allocator.free(got);
    try expectListEqual("ref", got, &[_]f64{ nan(), nan(), 1.0, 2.0, 3.0 });

    const backward = try formula.refValue(allocator, &x, -1);
    defer allocator.free(backward);
    try expectListEqual("ref(-1)", backward, &[_]f64{ 2.0, 3.0, 4.0, 5.0, nan() });
}

test "ma 完整窗口才计算" {
    const allocator = std.testing.allocator;
    const x = [_]f64{ 1.0, 2.0, 3.0, 4.0, 5.0 };
    const got = try formula.ma(allocator, &x, 3);
    defer allocator.free(got);
    try expectListEqual("ma", got, &[_]f64{ nan(), nan(), 2.0, 3.0, 4.0 });

    const too_wide = try formula.ma(allocator, &x, 9);
    defer allocator.free(too_wide);
    try expectListEqual("ma(9)", too_wide, &[_]f64{ nan(), nan(), nan(), nan(), nan() });
}

test "ema 首值等于源序列首值" {
    const allocator = std.testing.allocator;
    const x = [_]f64{ 1.0, 2.0, 3.0, 4.0 };
    const got = try formula.ema(allocator, &x, 4);
    defer allocator.free(got);
    try expectListEqual("ema", got, &[_]f64{ 1.0, 1.4, 2.04, 2.824 });
}

test "sma 中国式平滑" {
    const allocator = std.testing.allocator;
    const x = [_]f64{ 1.0, 2.0, 3.0 };
    const got = try formula.sma(allocator, &x, 3, 1);
    defer allocator.free(got);
    try expectListEqual("sma", got, &[_]f64{ 1.0, 1.3333333333333335, 1.888888888888889 });
}

test "滚动族 hhv/llv/sum/std" {
    const allocator = std.testing.allocator;
    const x = [_]f64{ 1.0, 3.0, 2.0, 5.0, 4.0 };

    const high = try formula.hhv(allocator, &x, 3);
    defer allocator.free(high);
    try expectListEqual("hhv", high, &[_]f64{ nan(), nan(), 3.0, 5.0, 5.0 });

    const low = try formula.llv(allocator, &x, 3);
    defer allocator.free(low);
    try expectListEqual("llv", low, &[_]f64{ nan(), nan(), 1.0, 2.0, 2.0 });

    const total = try formula.sum(allocator, &x, 2);
    defer allocator.free(total);
    try expectListEqual("sum", total, &[_]f64{ nan(), 4.0, 5.0, 7.0, 9.0 });

    const deviation = try formula.stddev(allocator, &x, 3);
    defer allocator.free(deviation);
    try expectListEqual("std", deviation, &[_]f64{ nan(), nan(), 0.816496580927726, 1.247219128924647, 1.247219128924647 });
}

test "wma 与 sqrt" {
    const allocator = std.testing.allocator;
    const x = [_]f64{ 1.0, 2.0, 3.0, 4.0 };
    const weighted = try formula.wma(allocator, &x, 3);
    defer allocator.free(weighted);
    try expectListEqual("wma", weighted, &[_]f64{ nan(), nan(), 2.3333333333333335, 3.3333333333333335 });

    const roots = try formula.sqrt(allocator, &[_]f64{ 1.0, 4.0, 9.0, 16.0 });
    defer allocator.free(roots);
    try expectListEqual("sqrt", roots, &[_]f64{ 1.0, 2.0, 3.0, 4.0 });
}

test "条件族 count/filter/cross" {
    const allocator = std.testing.allocator;
    const condition = [_]f64{ 1.0, 0.0, 1.0, 1.0, 0.0, 1.0 };

    const total = try formula.count(allocator, &condition, 3);
    defer allocator.free(total);
    try expectListEqual("count", total, &[_]f64{ 0.0, 0.0, 2.0, 2.0, 2.0, 2.0 });

    const gated = try formula.filter(allocator, &condition, 2);
    defer allocator.free(gated);
    try expectListEqual("filter", gated, &[_]f64{ 1.0, 0.0, 0.0, 1.0, 0.0, 0.0 });

    const s1 = [_]f64{ 1.0, 2.0, 1.0, 3.0 };
    const s2 = [_]f64{ 1.5, 1.5, 2.5, 2.5 };
    const flags = try crossAsFloat(allocator, &s1, &s2);
    defer allocator.free(flags);
    try expectListEqual("cross", flags, &[_]f64{ 0.0, 1.0, 0.0, 1.0 });
}

test "极值位置与 dma/max/min" {
    const allocator = std.testing.allocator;
    const x = [_]f64{ 1.0, 5.0, 3.0, 5.0, 2.0 };

    const high = try formula.hhvbars(allocator, &x, 3);
    defer allocator.free(high);
    try expectListEqual("hhvbars", high, &[_]f64{ nan(), nan(), 1.0, 0.0, 1.0 });

    const low = try formula.llvbars(allocator, &x, 3);
    defer allocator.free(low);
    try expectListEqual("llvbars", low, &[_]f64{ nan(), nan(), 2.0, 1.0, 0.0 });

    const source = [_]f64{ 1.0, 2.0, 3.0, 4.0 };
    const alpha = [_]f64{ 2.0, 2.0, 2.0, 2.0 };
    const average = try formula.dma(allocator, &source, &alpha);
    defer allocator.free(average);
    try expectListEqual("dma", average, &[_]f64{ 1.0, 1.5, 2.25, 3.125 });

    const left = [_]f64{ 1.0, 5.0, 3.0 };
    const right = [_]f64{ 2.0, 4.0, 6.0 };
    const upper = try formula.maxSeries(allocator, &left, &right);
    defer allocator.free(upper);
    try expectListEqual("max", upper, &[_]f64{ 2.0, 5.0, 6.0 });

    const lower = try formula.minSeries(allocator, &left, &right);
    defer allocator.free(lower);
    try expectListEqual("min", lower, &[_]f64{ 1.0, 4.0, 3.0 });
}

test "比较族 gt/gte/lt/lte/eq/neq 与逻辑族" {
    const allocator = std.testing.allocator;
    const x = [_]f64{ 1.0, 2.0, 3.0, 4.0 };
    const y = [_]f64{ 2.0, 2.0, 2.0, 2.0 };

    const greater = try compareAsFloat(formula.gt, allocator, &x, &y);
    defer allocator.free(greater);
    try expectListEqual("gt", greater, &[_]f64{ 0.0, 0.0, 1.0, 1.0 });

    const equal = try compareAsFloat(formula.eq, allocator, &x, &y);
    defer allocator.free(equal);
    try expectListEqual("eq", equal, &[_]f64{ 0.0, 1.0, 0.0, 0.0 });

    const left = [_]f64{ 1.0, 1.0, 0.0, 0.0 };
    const right = [_]f64{ 1.0, 0.0, 1.0, 0.0 };
    const conjunction = try logicBinaryAsFloat(formula.logicAnd, allocator, &left, &right);
    defer allocator.free(conjunction);
    try expectListEqual("and", conjunction, &[_]f64{ 1.0, 0.0, 0.0, 0.0 });

    const disjunction = try logicBinaryAsFloat(formula.logicOr, allocator, &left, &right);
    defer allocator.free(disjunction);
    try expectListEqual("or", disjunction, &[_]f64{ 1.0, 1.0, 1.0, 0.0 });

    const negation = try compareNot(allocator, &left);
    defer allocator.free(negation);
    try expectListEqual("not", negation, &[_]f64{ 0.0, 0.0, 1.0, 1.0 });
}

test "条件选择 if/ifn" {
    const allocator = std.testing.allocator;
    const condition = [_]f64{ 1.0, 0.0, 1.0, 1.0, 0.0, 1.0 };
    const values = [_]f64{ 10.0, 20.0, 30.0, 40.0, 50.0, 60.0 };

    const chosen = try ifBranch(allocator, &condition, &values, -1.0, false);
    defer allocator.free(chosen);
    try expectListEqual("if", chosen, &[_]f64{ 10.0, -1.0, 30.0, 40.0, -1.0, 60.0 });

    const inverted = try ifBranch(allocator, &condition, &values, -1.0, true);
    defer allocator.free(inverted);
    try expectListEqual("ifn", inverted, &[_]f64{ -1.0, 20.0, -1.0, -1.0, 50.0, -1.0 });
}

test "条件距离族 barslast/barslastcount/barslasts/barssincen" {
    const allocator = std.testing.allocator;
    const condition = [_]f64{ 1.0, 0.0, 1.0, 1.0, 0.0, 1.0 };

    const distance = try formula.barslast(allocator, &condition);
    defer allocator.free(distance);
    try expectListEqual("barslast", distance, &[_]f64{ 0.0, 1.0, 0.0, 0.0, 1.0, 0.0 });

    const streak = try formula.barslastcount(allocator, &condition);
    defer allocator.free(streak);
    try expectListEqual("barslastcount", streak, &[_]f64{ 1.0, 0.0, 1.0, 2.0, 0.0, 1.0 });

    const nth = try formula.barslasts(allocator, &condition, 2);
    defer allocator.free(nth);
    try expectListEqual("barslasts", nth, &[_]f64{ nan(), nan(), 2.0, 1.0, 2.0, 2.0 });

    const since = try formula.barssincen(allocator, &condition, 3);
    defer allocator.free(since);
    try expectListEqual("barssincen", since, &[_]f64{ 0.0, 0.0, 2.0, 1.0, 2.0, 2.0 });
}

test "last 窗口不足为 NaN" {
    const allocator = std.testing.allocator;
    const condition = [_]f64{ 1.0, 1.0, 1.0, 0.0, 1.0, 1.0 };
    const result = try formula.last(allocator, &condition, 3, 1);
    defer allocator.free(result);
    try expectListEqual("last", result, &[_]f64{ nan(), nan(), nan(), 1.0, 0.0, 0.0 });
}

test "fixture cases.csv 逐行对拍" {
    const allocator = std.testing.allocator;
    const text = try readFixture(allocator);
    defer allocator.free(text);

    var lines = std.mem.splitScalar(u8, text, '\n');
    var checked: usize = 0;
    var header_seen = false;
    while (lines.next()) |raw| {
        const line = std.mem.trimEnd(u8, raw, "\r");
        if (line.len == 0) continue;
        if (!header_seen) {
            try std.testing.expectEqualStrings("formula,x,y,n,m,expected", line);
            header_seen = true;
            continue;
        }

        var fields: [6][]const u8 = undefined;
        var count: usize = 0;
        var cells = std.mem.splitScalar(u8, line, ',');
        while (cells.next()) |cell| {
            if (count == fields.len) return error.TooManyColumns;
            fields[count] = cell;
            count += 1;
        }
        if (count != fields.len) return error.TooFewColumns;

        const name = fields[0];
        const x = try parseList(allocator, fields[1]);
        defer allocator.free(x);
        const y = try parseList(allocator, fields[2]);
        defer allocator.free(y);
        const n = try parsePeriod(fields[3]);
        const m = try parsePeriod(fields[4]);
        const want = try parseList(allocator, fields[5]);
        defer allocator.free(want);

        const got = if (std.mem.eql(u8, name, "abs"))
            try formula.abs(allocator, x)
        else if (std.mem.eql(u8, name, "ref"))
            try formula.refValue(allocator, x, n)
        else if (std.mem.eql(u8, name, "ma"))
            try formula.ma(allocator, x, n)
        else if (std.mem.eql(u8, name, "ema"))
            try formula.ema(allocator, x, n)
        else if (std.mem.eql(u8, name, "sma"))
            try formula.sma(allocator, x, n, m)
        else if (std.mem.eql(u8, name, "hhv"))
            try formula.hhv(allocator, x, n)
        else if (std.mem.eql(u8, name, "llv"))
            try formula.llv(allocator, x, n)
        else if (std.mem.eql(u8, name, "sum"))
            try formula.sum(allocator, x, n)
        else if (std.mem.eql(u8, name, "std"))
            try formula.stddev(allocator, x, n)
        else if (std.mem.eql(u8, name, "wma"))
            try formula.wma(allocator, x, n)
        else if (std.mem.eql(u8, name, "sqrt"))
            try formula.sqrt(allocator, x)
        else if (std.mem.eql(u8, name, "count"))
            try formula.count(allocator, x, n)
        else if (std.mem.eql(u8, name, "filter"))
            try formula.filter(allocator, x, n)
        else if (std.mem.eql(u8, name, "hhvbars"))
            try formula.hhvbars(allocator, x, n)
        else if (std.mem.eql(u8, name, "llvbars"))
            try formula.llvbars(allocator, x, n)
        else if (std.mem.eql(u8, name, "cross"))
            try crossAsFloat(allocator, x, y)
        else if (std.mem.eql(u8, name, "dma"))
            try formula.dma(allocator, x, y)
        else if (std.mem.eql(u8, name, "max"))
            try formula.maxSeries(allocator, x, y)
        else if (std.mem.eql(u8, name, "min"))
            try formula.minSeries(allocator, x, y)
        else if (std.mem.eql(u8, name, "gt"))
            try compareAsFloat(formula.gt, allocator, x, y)
        else if (std.mem.eql(u8, name, "gte"))
            try compareAsFloat(formula.gte, allocator, x, y)
        else if (std.mem.eql(u8, name, "lt"))
            try compareAsFloat(formula.lt, allocator, x, y)
        else if (std.mem.eql(u8, name, "lte"))
            try compareAsFloat(formula.lte, allocator, x, y)
        else if (std.mem.eql(u8, name, "eq"))
            try compareAsFloat(formula.eq, allocator, x, y)
        else if (std.mem.eql(u8, name, "neq"))
            try compareAsFloat(formula.neq, allocator, x, y)
        else if (std.mem.eql(u8, name, "and"))
            try logicBinaryAsFloat(formula.logicAnd, allocator, x, y)
        else if (std.mem.eql(u8, name, "or"))
            try logicBinaryAsFloat(formula.logicOr, allocator, x, y)
        else if (std.mem.eql(u8, name, "not"))
            try compareNot(allocator, x)
        else if (std.mem.eql(u8, name, "barslast"))
            try formula.barslast(allocator, x)
        else if (std.mem.eql(u8, name, "barslastcount"))
            try formula.barslastcount(allocator, x)
        else if (std.mem.eql(u8, name, "barslasts"))
            try formula.barslasts(allocator, x, n)
        else if (std.mem.eql(u8, name, "barssincen"))
            try formula.barssincen(allocator, x, n)
        else if (std.mem.eql(u8, name, "if"))
            try ifBranch(allocator, x, y, @floatFromInt(n), false)
        else if (std.mem.eql(u8, name, "ifn"))
            try ifBranch(allocator, x, y, @floatFromInt(n), true)
        else if (std.mem.eql(u8, name, "last"))
            try formula.last(allocator, x, n, m)
        else
            return error.UnknownFormula;
        defer allocator.free(got);
        try expectListEqual(name, got, want);
        checked += 1;
    }
    try std.testing.expect(checked >= 35);
}
