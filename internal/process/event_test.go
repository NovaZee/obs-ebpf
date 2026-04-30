package process

import (
	"encoding/binary"
	"testing"
)

func TestFormatExecEvent(t *testing.T) {
	event := ExecEvent{
		PID:      1234,
		PPID:     1000,
		UID:      501,
		Comm:     "bash",
		Filename: "/usr/bin/ls",
	}

	got := FormatExecEvent(event)
	want := "pid=1234 ppid=1000 uid=501 comm=bash filename=/usr/bin/ls"

	if got != want {
		t.Fatalf("期望格式化结果为 %q，实际为 %q", want, got)
	}
}

func TestFormatExecEventUsesPlaceholderForEmptyFilename(t *testing.T) {
	event := ExecEvent{
		PID:  42,
		PPID: 1,
		UID:  0,
		Comm: "sh",
	}

	got := FormatExecEvent(event)
	want := "pid=42 ppid=1 uid=0 comm=sh filename=<unknown>"

	if got != want {
		t.Fatalf("期望空文件名使用占位符，实际为 %q", got)
	}
}

func TestDecodeExecEvent(t *testing.T) {
	raw := make([]byte, RawExecEventSize)
	binary.LittleEndian.PutUint32(raw[0:4], 1234)
	binary.LittleEndian.PutUint32(raw[4:8], 1000)
	binary.LittleEndian.PutUint32(raw[8:12], 501)
	copy(raw[12:28], []byte("bash\x00"))
	copy(raw[28:284], []byte("/usr/bin/ls\x00"))

	event, err := DecodeExecEvent(raw)

	if err != nil {
		t.Fatalf("期望解码成功，实际错误为 %v", err)
	}
	if event.PID != 1234 || event.PPID != 1000 || event.UID != 501 {
		t.Fatalf("基础字段解码错误：%+v", event)
	}
	if event.Comm != "bash" {
		t.Fatalf("期望 comm 为 bash，实际为 %q", event.Comm)
	}
	if event.Filename != "/usr/bin/ls" {
		t.Fatalf("期望 filename 为 /usr/bin/ls，实际为 %q", event.Filename)
	}
}

func TestDecodeExecEventRejectsShortRecord(t *testing.T) {
	_, err := DecodeExecEvent(make([]byte, RawExecEventSize-1))

	if err == nil {
		t.Fatal("期望短记录返回错误，实际没有错误")
	}
}
