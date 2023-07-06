package config

var DefaultConfig = Config{
	TickRate: 20,
	Theme:    "simple",
}

type Config struct {
	TickRate int
	Theme    string
}

func NewConfig(tickRate int, theme string) (Config, error) {
	c := &Config{
		TickRate: tickRate,
		Theme:    theme,
	}

	if err := Validate(c); err != nil {
		return Config{}, err
	}

	return *c, nil
}
