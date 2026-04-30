# 进程观测示例

本目录用于学习 eBPF 进程观测。

## 阶段 1 目标

- 观测进程启动事件。
- 输出 PID、父 PID、UID、进程名。
- 尝试读取 exec 参数。
- 理解 tracepoint 或 kprobe 的挂载方式。
- 使用 Go 读取内核态事件。

## 当前文件

- `execsnoop.bpf.c`：内核态 eBPF 程序，挂载 `tracepoint/syscalls/sys_enter_execve`。
- `execsnoop.go`：Linux 下的示例入口，后续接入 Go loader。
- `execsnoop_unsupported.go`：非 Linux 下的提示入口。

## 学习重点

`execsnoop.bpf.c` 先关注 4 件事：

1. `SEC("tracepoint/syscalls/sys_enter_execve")` 表示程序挂载到 execve 系统调用入口。
2. `struct exec_event` 是内核态传给用户态的数据结构。
3. `BPF_MAP_TYPE_RINGBUF` 用来把事件从内核态送到用户态。
4. `bpf_probe_read_user_str` 从用户态地址读取执行文件路径。

## 运行

当前仓库在 macOS 上可以运行用户态测试：

```bash
go test ./...
go run ./examples/process
```

真实加载 eBPF 程序需要 Linux 环境：

```bash
cd examples/process
go generate .
go build -o ../../bin/execsnoop .
sudo ../../bin/execsnoop
```

也可以直接使用本目录的 Makefile：

```bash
make build
sudo ../../bin/execsnoop
```

运行后在另一个终端执行命令，例如：

```bash
/bin/ls
```

`execsnoop` 会输出 exec 事件：

```text
pid=1234 ppid=1000 uid=0 comm=ls filename=/bin/ls
```
