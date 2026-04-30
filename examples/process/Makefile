BPF_CLANG ?= clang

.PHONY: generate build test clean

generate:
	go generate .

build: generate
	mkdir -p ../../bin
	go build -o ../../bin/execsnoop .

test:
	go test ./...

clean:
	rm -f execsnoop_bpfeb.go execsnoop_bpfeb.o execsnoop_bpfel.go execsnoop_bpfel.o
