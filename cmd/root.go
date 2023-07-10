package cmd

import (
	"flag"
	"os"

	"github.com/uniyuni1101/retroclock/config"
	"github.com/uniyuni1101/retroclock/engine"
	"github.com/uniyuni1101/retroclock/render"
)

var RootCmd = Command{
	config:        &config.Config{},
	defaultConfig: &config.DefaultConfig,
	flagset:       flag.NewFlagSet("retroclock", flag.ExitOnError),
}

type Command struct {
	config        *config.Config
	defaultConfig *config.Config
	flagset       *flag.FlagSet
}

func (c *Command) init() error {
	c.flagset.IntVar(&c.config.TickRate, "rate", c.defaultConfig.TickRate, "display update rate")
	c.flagset.IntVar(&c.config.TickRate, "r", c.defaultConfig.TickRate, "display update rate (short)")

	c.flagset.StringVar(&c.config.Theme, "theme", c.defaultConfig.Theme, "display theme setting")
	c.flagset.StringVar(&c.config.Theme, "t", c.defaultConfig.Theme, "display theme setting (short)")

	c.flagset.Parse(os.Args[1:])
	return nil
}

func (c *Command) validate() error {
	return config.Validate(c.config)
}

func (c *Command) Execute() error {
	if err := c.init(); err != nil {
		return err
	}
	if err := c.validate(); err != nil {
		return err
	}

	ticker := engine.NewTickController(engine.TickRate(c.config.TickRate))
	render := render.NewSimple(os.Stdout)

	engine := engine.Engine{
		Config: *c.config,
		Ticker: ticker,
		Render: render,
	}

	engine.Run()

	return nil
}
