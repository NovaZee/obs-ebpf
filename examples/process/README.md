# 进程观测示例

本目录用于学习 eBPF 进程观测。

第一批目标：

- 观测进程启动事件。
- 输出 PID、父 PID、UID、进程名。
- 尝试读取 exec 参数。
- 理解 tracepoint 或 kprobe 的挂载方式。
- 使用 Go 读取内核态事件。

后续这里会放置对应的 Go loader 和 C eBPF 程序。
