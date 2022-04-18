package main

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"path/filepath"

	"github.com/nxadm/tail"
)

type Color string

const (
	Reset  Color = "\033[0m"
	Red    Color = "\033[31m"
	Green  Color = "\033[32m"
	Yellow Color = "\033[33m"
	Purple Color = "\033[35m"
	Cyan   Color = "\033[36m"
)

const n = 5

var colors = [n]Color{Red, Green, Yellow, Purple, Cyan}

func (c Color) String() string {
	return string(c)
}

func randColor() Color {
	idx := rand.Intn(n)

	return colors[idx]
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
		err := tailf(path, randColor(), config, 10)
		if err != nil {
			panic(err)
		}
	}

	ch := make(chan bool)
	<-ch
}

func tailf(path string, color Color, config tail.Config, last int) error {

	t, err := tail.TailFile(
		path,
		config,
	)

	if err != nil {
		return err
	}

	base := filepath.Base(path)

	go func(tag string) {
		for line := range t.Lines {
			fmt.Printf("%s---> [%s]: %s\n", color, tag, line.Text)
		}
	}(base)

	return nil
}
