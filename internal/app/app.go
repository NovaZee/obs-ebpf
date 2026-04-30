package app

import (
	"fmt"
	"io"
	"strings"
)

const overview = `obs-ebpf 是一个 Go 优先的 eBPF 可观测性学习与实践项目。

当前优先方向：
  process  进程观测学习阶段
  network  网络观测学习阶段

运行示例：
  obs-ebpf process
  obs-ebpf network
`

// Run 执行命令并返回进程退出码，便于 main 和测试复用同一套行为。
func Run(args []string, stdout io.Writer, stderr io.Writer) int {
	if len(args) <= 1 {
		fmt.Fprint(stdout, overview)
		return 0
	}

	switch strings.ToLower(args[1]) {
	case "process":
		fmt.Fprintln(stdout, "进程观测学习阶段：先从 exec/exit 事件开始，打通 eBPF 到 Go 的事件链路。")
		return 0
	case "network":
		fmt.Fprintln(stdout, "网络观测学习阶段：先从 TCP connect 事件开始，采集 PID、进程名、目标地址和端口。")
		return 0
	case "help", "-h", "--help":
		fmt.Fprint(stdout, overview)
		return 0
	default:
		fmt.Fprintf(stderr, "未知命令：%s\n\n", args[1])
		fmt.Fprint(stderr, overview)
		return 2
	}
}
