package config

import "errors"

const (
	MinTickPerSecond = 1
	MaxTickPerSecond = 120
)

var (
	ErrOutOfRangeTickPerSecond = errors.New("tick per second value out of range, please choose a value within the range of 1 to 120")
	ErrNotFoundTheme           = errors.New("not found theme")
)

func Validate(cfg *Config) error {

	if err := validateTickPerSecond(cfg); err != nil {
		return err
	}

	if err := validateTheme(cfg); err != nil {
		return err
	}

	return nil
}

func validateTickPerSecond(cfg *Config) error {
	if cfg.TickPerSecond < MinTickPerSecond || MaxTickPerSecond < cfg.TickPerSecond {
		return ErrOutOfRangeTickPerSecond
	}

	return nil
}

func validateTheme(cfg *Config) error {

	if cfg.Theme != "simple" {
		return ErrNotFoundTheme
	}

	return nil
}
