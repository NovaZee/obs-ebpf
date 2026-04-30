package process

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

const (
	unknownFilename  = "<unknown>"
	commLen          = 16
	filenameLen      = 256
	RawExecEventSize = 12 + commLen + filenameLen
)

// ExecEvent 是进程 exec 观测的用户态事件模型。
type ExecEvent struct {
	PID      uint32
	PPID     uint32
	UID      uint32
	Comm     string
	Filename string
}

// FormatExecEvent 输出稳定的一行文本，便于学习阶段观察和测试。
func FormatExecEvent(event ExecEvent) string {
	filename := event.Filename
	if filename == "" {
		filename = unknownFilename
	}

	return fmt.Sprintf(
		"pid=%d ppid=%d uid=%d comm=%s filename=%s",
		event.PID,
		event.PPID,
		event.UID,
		event.Comm,
		filename,
	)
}

func DecodeExecEvent(raw []byte) (ExecEvent, error) {
	if len(raw) < RawExecEventSize {
		return ExecEvent{}, fmt.Errorf("exec event record too short: got %d want %d", len(raw), RawExecEventSize)
	}

	return ExecEvent{
		PID:      binary.LittleEndian.Uint32(raw[0:4]),
		PPID:     binary.LittleEndian.Uint32(raw[4:8]),
		UID:      binary.LittleEndian.Uint32(raw[8:12]),
		Comm:     cString(raw[12 : 12+commLen]),
		Filename: cString(raw[12+commLen : RawExecEventSize]),
	}, nil
}

func LearningHint() string {
	return "进程观测学习阶段：阶段 1 先实现 exec 事件，关注 PID、PPID、UID、进程名和执行文件路径。"
}

func cString(raw []byte) string {
	if index := bytes.IndexByte(raw, 0); index >= 0 {
		return string(raw[:index])
	}
	return string(raw)
}
