# obs-ebpf

`obs-ebpf` 是一个 Go 优先的 eBPF 可观测性学习与实践项目。

目标不是一开始就做一个复杂平台，而是先用小而清晰的示例打通 eBPF 可观测性的基础链路，再把成熟能力沉淀成一个可长期演进的命令行工具。

## 项目方向

- 学习 eBPF 在可观测性领域的核心能力。
- 编写进程观测、网络观测等方向的 eBPF 示例程序。
- 优先使用 Go 编写用户态加载器、CLI、事件读取和输出逻辑。
- 必要时使用 C 编写内核态 eBPF 程序，复杂处理尽量放在 Go 侧。

## 初始结构

```text
cmd/obs-ebpf/          主命令入口
docs/                  学习路线和项目文档
examples/process/      进程观测学习示例
examples/network/      网络观测学习示例
internal/app/          CLI 行为组织
internal/ebpf/         后续沉淀 eBPF 加载和公共能力
internal/output/       后续沉淀输出格式化能力
```

## 当前阶段

第一阶段优先学习和实践：

1. 进程观测：进程启动、退出、exec 参数、PID、UID、命令名。
2. 网络观测：TCP 连接、端口、远端地址、连接结果。

先让 `examples/` 中的小示例跑通，再把稳定代码沉淀到 `cmd/obs-ebpf` 和 `internal/`。

## 运行

```bash
go run ./cmd/obs-ebpf
go run ./cmd/obs-ebpf process
go run ./cmd/obs-ebpf network
```

## 验证

```bash
go test ./...
```
