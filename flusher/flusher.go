package flusher

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"text/template"

	"github.com/ainoya/sewer/drainer"
)

type Flusher struct {
	drainer  drainer.Drainer
	reader   *bufio.Reader
	template *template.Template
}

type MsgOutput struct {
	Message string
}

func NewFlusher(
	drainer drainer.Drainer,
	reader *bufio.Reader,
	templateText string,
) *Flusher {

	tmpl := template.Must(template.New("messageTemplate").Parse(templateText))

	return &Flusher{
		drainer:  drainer,
		reader:   reader,
		template: tmpl,
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
		message += input
	}
	var out bytes.Buffer
	f.template.Execute(&out, MsgOutput{Message: message})
	return f.drainer.Drain(out.String())
}
