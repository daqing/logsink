package main

import (
	"fmt"
	"io"
	"os"
	"os/signal"
	"path/filepath"
	"strings"

	"github.com/nxadm/tail"
)

type Color string

const (
	Reset  Color = "\033[0m"
	Red    Color = "\033[31m"
	Green  Color = "\033[32m"
	Yellow Color = "\033[33m"
	Blue   Color = "\033[34;1m"
	Purple Color = "\033[35m"
	Cyan   Color = "\033[36m"
	White  Color = "\033[37m"
)

const (
	SEP = string(filepath.Separator)
)

const nColors = 7

var colors = [nColors]Color{Green, Yellow, Blue, White, Purple, Cyan, Red}

// global index to assign color
var idx = 0

func (c Color) String() string {
	return string(c)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide paths to watch")
		return
	}

	config := tail.Config{
		Follow: true,
		ReOpen: true,
		Location: &tail.SeekInfo{
			Offset: 0,
			Whence: io.SeekEnd,
		},
	}

	for _, path := range os.Args[1:] {
		err := tailf(path, assignColor(), config, 10)
		if err != nil {
			panic(err)
		}
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c

	fmt.Println(Reset, "bye")
}

func assignColor() Color {
	if idx >= nColors {
		idx = 0
	}

	c := colors[idx]

	idx += 1

	return c
}

func tailf(path string, color Color, config tail.Config, last int) error {

	t, err := tail.TailFile(
		path,
		config,
	)

	if err != nil {
		return err
	}

	go func(tag string) {
		for line := range t.Lines {
			fmt.Printf("%s===> [%s]:%s %s\n", color, tag, Reset, line.Text)
		}
	}(pathTag(path))

	return nil
}

func pathTag(path string) string {
	parts := strings.Split(path, SEP)
	length := len(parts)

	if length < 3 {
		return path
	}

	return strings.Join(parts[length-3:length], SEP)
}
