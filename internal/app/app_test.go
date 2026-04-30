package app

import (
	"bytes"
	"testing"
)

func TestRunWithoutArgsPrintsOverview(t *testing.T) {
	var out bytes.Buffer

	code := Run([]string{"obs-ebpf"}, &out, &out)

	if code != 0 {
		t.Fatalf("期望退出码为 0，实际为 %d", code)
	}
	got := out.String()
	want := "obs-ebpf 是一个 Go 优先的 eBPF 可观测性学习与实践项目"
	if !bytes.Contains([]byte(got), []byte(want)) {
		t.Fatalf("期望输出包含 %q，实际输出为 %q", want, got)
	}
}

func TestRunProcessCommandPrintsLearningHint(t *testing.T) {
	var out bytes.Buffer

	code := Run([]string{"obs-ebpf", "process"}, &out, &out)

	if code != 0 {
		t.Fatalf("期望退出码为 0，实际为 %d", code)
	}
	got := out.String()
	want := "进程观测学习阶段"
	if !bytes.Contains([]byte(got), []byte(want)) {
		t.Fatalf("期望输出包含 %q，实际输出为 %q", want, got)
	}
}

func TestRunNetworkCommandPrintsLearningHint(t *testing.T) {
	var out bytes.Buffer

	code := Run([]string{"obs-ebpf", "network"}, &out, &out)

	if code != 0 {
		t.Fatalf("期望退出码为 0，实际为 %d", code)
	}
	got := out.String()
	want := "网络观测学习阶段"
	if !bytes.Contains([]byte(got), []byte(want)) {
		t.Fatalf("期望输出包含 %q，实际输出为 %q", want, got)
	}
}

func TestRunUnknownCommandReturnsError(t *testing.T) {
	var out bytes.Buffer

	code := Run([]string{"obs-ebpf", "unknown"}, &out, &out)

	if code != 2 {
		t.Fatalf("期望未知命令退出码为 2，实际为 %d", code)
	}
	got := out.String()
	want := "未知命令"
	if !bytes.Contains([]byte(got), []byte(want)) {
		t.Fatalf("期望输出包含 %q，实际输出为 %q", want, got)
	}
}
