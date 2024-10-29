package main

import (
	"flag"
	"gart/cmd"
	"gart/tui"
)

var (
	mode string
)

func init() {
	flag.StringVar(&mode, "mode", "tui", "运行模式，支持tui，cmd，默认tui模式。cmd为普通命令行，tui添加了UI效果")
}

func main() {
	flag.Parse()
	if mode == "tui" {
		tui.Execute()
	} else {
		cmd.Execute()
	}
}

// go build -ldflags='-s -w'
