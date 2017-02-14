package flusher

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
	"text/template"

	"github.com/ainoya/sewer/drainer"
)

type Flusher struct {
	drainers  []drainer.Drainer
	reader    *bufio.Reader
	template  *template.Template
	eachlines bool
}

type MsgOutput struct {
	Message string
}

func NewFlusher(
	drainers []drainer.Drainer,
	reader *bufio.Reader,
	templateText string,
	eachlines bool,
) *Flusher {

	tmpl := template.Must(template.New("messageTemplate").Parse(templateText))

	return &Flusher{
		drainers:  drainers,
		reader:    reader,
		template:  tmpl,
		eachlines: eachlines,
	}
}

func (f Flusher) Flush() error {
	message := ""
	var err error
	var out bytes.Buffer

	for {
		input, err := f.reader.ReadString('\n')
		if err != nil && err == io.EOF {
			break
		}

		fmt.Printf("%s", input)
		message += input
		if f.eachlines {
			for _, dnr := range f.drainers {
				var o bytes.Buffer
				f.template.Execute(&o, MsgOutput{Message: strings.Trim(input, "\n")})
				oStr := o.String()
				err = dnr.Drain(oStr)
			}
		}
	}

	if f.eachlines {
		return err
	}

	f.template.Execute(&out, MsgOutput{Message: message})
	var outStr = out.String()

	for _, dnr := range f.drainers {
		err = dnr.Drain(outStr)
	}
	return err
}
