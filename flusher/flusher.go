package flusher

import (
	"bufio"
	"fmt"
	"io"

	"github.com/ainoya/sewer/drainer"
)

type Flusher struct {
	drainer drainer.Drainer
	reader  *bufio.Reader
}

func NewFlusher(
	drainer drainer.Drainer,
	reader *bufio.Reader) *Flusher {
	return &Flusher{
		drainer: drainer,
		reader:  reader,
	}
}

func (f Flusher) Flush() error {
	message := ""
	for {
		input, err := f.reader.ReadString('\n')
		if err != nil && err == io.EOF {
			break
		}

		fmt.Printf("%s", input)
		message += input + "\n"
	}
	return f.drainer.Drain(message)
}
