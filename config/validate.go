package config

import "errors"

const (
	MinTickRate = 1
	MaxTickRate = 120
)

var (
	ErrOutOfRangeTickRate = errors.New("tick per second value out of range, please choose a value within the range of 1 to 120")
	ErrNotFoundTheme           = errors.New("not found theme")
)

func Validate(cfg *Config) error {

	if err := validateTickRate(cfg); err != nil {
		return err
	}

	if err := validateTheme(cfg); err != nil {
		return err
	}

	return nil
}

func validateTickRate(cfg *Config) error {
	if cfg.TickRate < MinTickRate || MaxTickRate < cfg.TickRate {
		return ErrOutOfRangeTickRate
	}

	return nil
}

func validateTheme(cfg *Config) error {

	if cfg.Theme != "simple" {
		return ErrNotFoundTheme
	}

	return nil
}
