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
	drainers []drainer.Drainer
	reader   *bufio.Reader
	template *template.Template
}

type MsgOutput struct {
	Message string
}

func NewFlusher(
	drainers []drainer.Drainer,
	reader *bufio.Reader,
	templateText string,
) *Flusher {

	tmpl := template.Must(template.New("messageTemplate").Parse(templateText))

	return &Flusher{
		drainers: drainers,
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
	var outStr = out.String()

	var err error
	for _, dnr := range f.drainers {
		err = dnr.Drain(outStr)
	}
	return err
}
