//go:build linux

package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/NovaZee/obs-ebpf/internal/process"
	"github.com/cilium/ebpf/link"
	"github.com/cilium/ebpf/perf"
	"github.com/cilium/ebpf/rlimit"
)

//go:generate go run github.com/cilium/ebpf/cmd/bpf2go -tags linux -target bpfel execsnoop execsnoop.bpf.c -- -I.

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	if err := rlimit.RemoveMemlock(); err != nil {
		return fmt.Errorf("remove memlock: %w", err)
	}

	var objects execsnoopObjects
	if err := loadExecsnoopObjects(&objects, nil); err != nil {
		return fmt.Errorf("load eBPF objects: %w", err)
	}
	defer objects.Close()

	tracepoint, err := link.Tracepoint("syscalls", "sys_enter_execve", objects.HandleExecve, nil)
	if err != nil {
		return fmt.Errorf("attach execve tracepoint: %w", err)
	}
	defer tracepoint.Close()

	reader, err := perf.NewReader(objects.Events, os.Getpagesize()*8)
	if err != nil {
		return fmt.Errorf("open perf event reader: %w", err)
	}
	defer reader.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-stop
		reader.Close()
	}()

	for {
		record, err := reader.Read()
		if err != nil {
			if errors.Is(err, perf.ErrClosed) {
				return nil
			}
			return fmt.Errorf("read perf event: %w", err)
		}
		if record.LostSamples != 0 {
			continue
		}

		event, err := process.DecodeExecEvent(record.RawSample)
		if err != nil {
			return fmt.Errorf("decode exec event: %w", err)
		}
		fmt.Fprintln(os.Stdout, process.FormatExecEvent(event))
	}
}
