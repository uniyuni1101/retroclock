package config

var DefaultConfig = Config{
	TickPerSecond: 20,
	Theme:         "simple",
}

type Config struct {
	TickPerSecond int
	Theme         string
}

func NewConfig(tickPerSecond int, theme string) (Config, error) {
	c := &Config{
		TickPerSecond: tickPerSecond,
		Theme:         theme,
	}

	if err := Validate(c); err != nil {
		return Config{}, err
	}

	return *c, nil
}
