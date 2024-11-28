package main

import (
	"log"
	"os"
	"time"

	"github.com/go-vgo/robotgo"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "wiggle",
		Usage: "wiggle thy mouse",
		Flags: []cli.Flag{
			&cli.DurationFlag{
				Name:    "pause",
				Value:   1 * time.Minute,
				Usage:   "time between mouse movements",
				Aliases: []string{"p"},
			},
		},
		Action: func(ctx *cli.Context) error {
			dur := ctx.Duration("pause")
			for {
				robotgo.MoveSmooth(200, 200, 1.0, 1.0)
				time.Sleep(dur)
				robotgo.MoveSmooth(500, 200, 1.0, 1.0)
				time.Sleep(dur)
			}
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
