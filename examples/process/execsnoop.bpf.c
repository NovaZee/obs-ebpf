//go:build ignore

#include "vmlinux.h"

char LICENSE[] SEC("license") = "Dual BSD/GPL";

#define TASK_COMM_LEN 16
#define FILENAME_LEN 256

struct exec_event {
	__u32 pid;
	__u32 ppid;
	__u32 uid;
	char comm[TASK_COMM_LEN];
	char filename[FILENAME_LEN];
};

struct {
	__uint(type, BPF_MAP_TYPE_PERF_EVENT_ARRAY);
	__uint(key_size, sizeof(__u32));
	__uint(value_size, sizeof(__u32));
} events SEC(".maps");

SEC("tracepoint/syscalls/sys_enter_execve")
int handle_execve(struct trace_event_raw_sys_enter *ctx)
{
	struct exec_event event = {};
	struct task_struct *task;
	const char *filename;
	__u64 pid_tgid;
	__u64 uid_gid;

	pid_tgid = bpf_get_current_pid_tgid();
	uid_gid = bpf_get_current_uid_gid();
	task = (struct task_struct *)bpf_get_current_task();
	filename = (const char *)ctx->args[0];

	event.pid = pid_tgid >> 32;
	event.uid = uid_gid & 0xffffffff;
	bpf_probe_read(&task, sizeof(task), &task->real_parent);
	bpf_probe_read(&event.ppid, sizeof(event.ppid), &task->tgid);
	bpf_get_current_comm(event.comm, sizeof(event.comm));
	bpf_probe_read_str(event.filename, sizeof(event.filename), filename);

	bpf_perf_event_output(ctx, &events, BPF_F_CURRENT_CPU, &event, sizeof(event));
	return 0;
}
