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
			Usage: "cartoon name path",
		},
		cli.StringFlag{
			Name:  "startchapter, sc",
			Usage: "number start chapter that want to download",
		},
		cli.StringFlag{
			Name:  "endchapter, ec",
			Usage: "number end chapter that want to download",
		},
		cli.StringFlag{
			Name:  "destination, d",
			Usage: "output destination path",
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

// ReadCLIParams ...
func (c *Config) ReadCLIParams() error {
	return c.CLI.Run(os.Args)

}
