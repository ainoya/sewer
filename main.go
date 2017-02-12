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
	}

	var d drainer.Drainer

	app.Action = func(c *cli.Context) error {
		info, _ := os.Stdin.Stat()

		if (info.Mode() & os.ModeCharDevice) == os.ModeCharDevice {
			fmt.Println("The command is intended to work with pipes.")
			fmt.Println("Usage:")
			fmt.Println("cat yourfile.txt | sewer --drain=github")
			return nil
		} else if info.Size() <= 0 {
			fmt.Println("Pipe input is null")
			return nil
		}

		reader := bufio.NewReader(os.Stdin)

		if c.String("drain") == "github" {
			var err error
			d, err = drainer.NewGitHubDrainer()
			if err != nil {
				fmt.Print(err)
				return nil
			}
		} else {
			return nil
		}

		f := flusher.NewFlusher(d, reader)
		f.Flush()

		return nil
	}

	app.Run(os.Args)
}
