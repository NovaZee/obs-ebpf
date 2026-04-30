package main

import (
	"os"

	"github.com/NovaZee/obs-ebpf/internal/app"
)

func main() {
	os.Exit(app.Run(os.Args, os.Stdout, os.Stderr))
}
