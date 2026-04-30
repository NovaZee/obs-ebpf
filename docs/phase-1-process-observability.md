# 阶段 1：进程观测

阶段 1 的目标是先完成 exec 事件观测的学习闭环。

## 当前进度

已经完成：

- 定义用户态 `ExecEvent` 模型。
- 定义稳定的一行文本输出格式。
- 在 `obs-ebpf process` 中输出阶段 1 学习提示。
- 添加 `examples/process/execsnoop.bpf.c` 作为内核态学习示例。
- 添加 Linux loader，使用 `cilium/ebpf` 挂载 execve tracepoint 并读取 ring buffer。

## 事件字段

第一版 exec 事件只关注这些字段：

- `pid`：当前进程 ID。
- `ppid`：父进程 ID。
- `uid`：发起 exec 的用户 ID。
- `comm`：内核中的进程名。
- `filename`：execve 的执行文件路径。

输出格式：

```text
pid=1234 ppid=1000 uid=501 comm=bash filename=/usr/bin/ls
```

## 为什么先做 exec

exec 是进程观测最适合入门的事件之一：

- 事件频率可控，容易触发和验证。
- 字段少，但覆盖 PID、UID、进程名、用户态字符串读取。
- 能自然引出 tracepoint、ring buffer、结构体对齐和 Go 侧事件解析。

## 下一步

Linux 环境编译运行：

```bash
cd examples/process
go generate .
go build -o ../../bin/execsnoop .
sudo ../../bin/execsnoop
```

下一步继续增强：

1. 解析 exec 参数。
2. 增加 exit 事件。
3. 增加过滤参数，例如按 UID、PID、进程名过滤。
4. 将 `obs-ebpf process` 接到真实 loader。
