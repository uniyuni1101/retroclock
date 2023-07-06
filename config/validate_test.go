package config

import (
	"testing"
)

func TestValidate_AllExecutionOfValidations(t *testing.T) {
	cases := map[string]struct {
		cfg  Config
		want error
	}{
		"valid all parameter": {
			Config{
				TickPerSecond: 20,
				Theme:         "simple",
			},
			nil,
		},
		"invalid TickPerSecond": {
			Config{
				TickPerSecond: 0,
				Theme:         "simple",
			},
			ErrOutOfRangeTickPerSecond,
		},
		"invalid Theme": {
			Config{
				TickPerSecond: 20,
				Theme:         "error",
			},
			ErrNotFoundTheme,
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			if err := Validate(&c.cfg); err != c.want {
				t.Errorf("got %v, want %v", err, c.want)
			}
		})
	}
}

func TestValidateTickPerSecond(t *testing.T) {
	cases := map[string]struct {
		cfg  Config
		want error
	}{
		"valid value between 1 and 120 when value is 1": {
			Config{TickPerSecond: 1}, nil,
		},
		"valid value between 1 and 120 when value is 120": {
			Config{TickPerSecond: 120}, nil,
		},
		"invalid value less than or equal to 0": {
			Config{TickPerSecond: 0},
			ErrOutOfRangeTickPerSecond,
		},
		"invalid value greater than 120": {
			Config{TickPerSecond: 121},
			ErrOutOfRangeTickPerSecond,
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			if err := validateTickPerSecond(&c.cfg); err != c.want {
				t.Errorf("got %v, want %v, config: %v", err, c.want, c.cfg)
			}
		})
	}
}

func TestValidateTheme(t *testing.T) {
	cases := map[string]struct {
		cfg  Config
		want error
	}{
		"valid theme find theme name": {
			Config{Theme: "simple"}, nil,
		},
		"invalid theme not found theme name": {
			Config{Theme: "pls_error"}, ErrNotFoundTheme,
		},
		// Todo: Add Nameing validation test
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			if err := validateTheme(&c.cfg); err != c.want {
				t.Errorf("got %v, want %v, config: %v", err, c.want, c.cfg)
			}
		})
	}
}
