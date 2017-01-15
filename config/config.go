package config

import (
	"os"

	"gopkg.in/urfave/cli.v1"
)

// Config ....
type Config struct {
	CLI            *cli.App
	CartoonProfile *CartoonProfile
}

// CartoonProfile ...
type CartoonProfile struct {
	DestinationPath string
	CartoonPath     string
	StartChapter    int
	EndChapter      int
}

// New ...
func New() *Config {
	config := &Config{}
	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "path, p",
			Value: "naruto",
			Usage: "input cartoon path",
		},
		cli.StringFlag{
			Name:  "startchapter, sc",
			Value: "1",
			Usage: "input start chapter that want to download",
		},
		cli.StringFlag{
			Name:  "endchapter, ec",
			Value: "12",
			Usage: "input end chapter that want to download",
		},
		cli.StringFlag{
			Name:  "destination, d",
			Value: "../mycartoon",
			Usage: "input destination path",
		},
	}

	app.Action = func(c *cli.Context) error {
		config.CartoonProfile = &CartoonProfile{
			DestinationPath: c.String("destination"),
			CartoonPath:     c.String("path"),
			StartChapter:    c.Int("startchapter"),
			EndChapter:      c.Int("endchapter"),
		}
		return nil
	}

	config.CLI = app
	return config
}

// Run ...
func (c *Config) ReadCLIParams() error {
	return c.CLI.Run(os.Args)

}
