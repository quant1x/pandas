# tests —— 跨语言 fixture 与一键验收

本目录放两套 fixture 与一个验收脚本。fixture 都是「输入 + 期望同表」，由 Python 生成器产出，
Go / C++ / Rust / Zig 各自读取、重算、逐值对拍；生成器固定输入、无随机数，重复执行逐字节一致。

| 目录 | 内容 | 入口 |
| --- | --- | --- |
| `formula/` | 通达信公式族 39 个公式，四端同值 | `go test ./formula/`、ctest、`cargo test`、`zig test` |
| `features/` | M8 特征（11 列 + `pred` + `pred_mlp`）与 ONNX 导出 | `go test ./tests/features/`、ctest、`cargo test` |

## 一键验收

```bash
python tests/run_all.py            # 生成器幂等 + Go / C++ / Rust / Zig
python tests/run_all.py --only go  # 只跑名字含 go 的步骤
```

`fixtures 幂等` 这一步会把两个生成器重跑一遍，比较产物的 SHA256：不一致即失败
（说明有人改了 fixture 没重跑生成器，或生成器引入不确定性）。

前置：C++ 需先 `cmake -S cpp -B cpp/build && cmake --build cpp/build --config Debug`；
其余只需对应工具链在 `PATH` 上。

## 新增 fixture

1. 在 `gen_*.py` 里追加用例并重跑生成器；
2. 四端各补一个分支（Go / C++ / Rust / Zig）；
3. `python tests/run_all.py` 全绿后提交。

期望值由 Python 侧生成；与某一端口径不一致时先在 issue 里记录，不直接改 fixture 迁就实现。
