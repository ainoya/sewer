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
		cli.StringFlag{
			Name:  "drain",
			Value: "github",
			Usage: "Destination where you want to notify piped messages",
		},
		cli.StringFlag{
			Name:  "template",
			Value: "{{ .Message }}",
			Usage: "Template format (STDIN is expanded as {{ .Message }} in the template)",
		},
	}

	var d drainer.Drainer

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
		fmt.Println(tmpl)

		drain := c.String("drain")
		var err error
		if drain == "github" {
			d, err = drainer.NewGitHubDrainer()
			if err != nil {
				fmt.Print(err)
				return nil
			}
		} else if drain == "slack" {
			d, err = drainer.NewSlackDrainer()
			if err != nil {
				fmt.Print(err)
				return nil
			}
		} else {
			return nil
		}

		f := flusher.NewFlusher(d, reader, tmpl)
		err = f.Flush()

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		return nil
	}

	app.Run(os.Args)
}
