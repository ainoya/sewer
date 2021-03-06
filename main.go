package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/ainoya/sewer/drainer"
	"github.com/ainoya/sewer/flusher"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.StringSliceFlag{
			Name:  "drain",
			Value: &cli.StringSlice{},
			Usage: "Destination where you want to notify piped messages",
		},
		cli.StringFlag{
			Name:  "template",
			Value: "{{ .Message }}",
			Usage: "Template format (STDIN is expanded as {{ .Message }} in the template)",
		},
		cli.BoolFlag{
			Name:   "eachlines",
			Hidden: false,
			Usage:  "If true, sewer sends message each lines.",
		},
	}

	app.Action = func(c *cli.Context) error {
		info, _ := os.Stdin.Stat()

		if (info.Mode() & os.ModeCharDevice) == os.ModeCharDevice {
			fmt.Println("The command is intended to work with pipes.")
			fmt.Println("Usage:")
			fmt.Println("cat yourfile.txt | sewer --drain=github")
			return nil
		}

		reader := bufio.NewReader(os.Stdin)
		tmpl := c.String("template")

		drains := c.StringSlice("drain")
		drainers, err := newDrainers(drains)

		if err != nil {
			fmt.Println(err)
			return nil
		}

		eachlines := c.Bool("eachlines")

		f := flusher.NewFlusher(drainers, reader, tmpl, eachlines)
		err = f.Flush()

		if err != nil {
			fmt.Println(err)
		}

		return nil
	}

	app.Run(os.Args)
}

func newDrainers(drains []string) ([]drainer.Drainer, error) {
	var drainers []drainer.Drainer
	for _, drain := range drains {
		var dnr drainer.Drainer
		var err error
		switch drain {
		case "github":
			dnr, err = drainer.NewGitHubDrainer()
		case "slack":
			dnr, err = drainer.NewSlackDrainer()
		default:
			err = fmt.Errorf("drain type %s is not defined.", drain)
		}
		if err != nil {
			return nil, err
		}

		drainers = append(drainers, dnr)
	}

	return drainers, nil
}
