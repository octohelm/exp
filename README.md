# exp

`exp` 是一组面向 Go 的实验性可复用能力，当前主要覆盖函数式组合、channel 观察模型与 iterator 变换。

仓库保持轻量：这里放的是可直接阅读和测试的小型包，而不是完整应用或多模块 workspace。

## 职责与边界

- 提供实验性的 Go library 包，包括 `fp`、`xchan`、`xiter`。
- 统一暴露仓库级的依赖整理、格式化与测试入口。
- 不负责应用运行入口、发布编排或多模块聚合。

## 主要入口

- [`fp`](/Users/morlay/src/github.com/octohelm/exp/fp) 用于函数式组合与管道式调用实验。
- [`xchan`](/Users/morlay/src/github.com/octohelm/exp/xchan) 用于基于 channel 的 observable / observer 抽象。
- [`xiter`](/Users/morlay/src/github.com/octohelm/exp/xiter) 用于 `iter.Seq` 相关的变换、聚合与时间类操作实验。
- [`justfile`](/Users/morlay/src/github.com/octohelm/exp/justfile) 提供仓库级稳定命令入口。
- [`AGENTS.md`](/Users/morlay/src/github.com/octohelm/exp/AGENTS.md) 定义本仓库协作约束与暂停条件。
