#ifndef __VMLINUX_H__
#define __VMLINUX_H__

typedef unsigned char __u8;
typedef short int __s16;
typedef short unsigned int __u16;
typedef int __s32;
typedef unsigned int __u32;
typedef long long int __s64;
typedef unsigned long long __u64;
typedef __u16 __le16;
typedef __u16 __be16;
typedef __u32 __be32;
typedef __u64 __be64;
typedef __u32 __wsum;
typedef __u16 __sum16;

#define SEC(name) __attribute__((section(name), used))
#define __uint(name, val) int (*name)[val]

enum bpf_map_type {
	BPF_MAP_TYPE_PERF_EVENT_ARRAY = 4,
};

#define BPF_F_INDEX_MASK 0xffffffffULL
#define BPF_F_CURRENT_CPU BPF_F_INDEX_MASK

static __u64 (*bpf_get_current_pid_tgid)(void) = (void *)14;
static __u64 (*bpf_get_current_uid_gid)(void) = (void *)15;
static long (*bpf_get_current_comm)(void *buf, __u32 size_of_buf) = (void *)16;
static long (*bpf_perf_event_output)(void *ctx, void *map, __u64 flags, void *data, __u64 size) = (void *)25;
static __u64 (*bpf_get_current_task)(void) = (void *)35;
static long (*bpf_probe_read)(void *dst, __u32 size, const void *unsafe_ptr) = (void *)4;
static long (*bpf_probe_read_str)(void *dst, __u32 size, const void *unsafe_ptr) = (void *)45;

struct task_struct {
	struct task_struct *real_parent;
	int tgid;
} __attribute__((preserve_access_index));

struct trace_entry {
	short unsigned int type;
	unsigned char flags;
	unsigned char preempt_count;
	int pid;
};

struct trace_event_raw_sys_enter {
	struct trace_entry ent;
	long int id;
	unsigned long int args[6];
	char __data[0];
};

#endif
