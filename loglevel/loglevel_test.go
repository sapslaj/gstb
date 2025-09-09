package loglevel

import (
	"errors"
	"log/slog"
	"strings"
	"testing"
)

func TestParseLogLevel(t *testing.T) {
	tests := map[string]struct {
		input            string
		expect           slog.Level
		errorIs          error
		errorMsgContains string
	}{
		"default info": {
			input:  "",
			expect: slog.LevelInfo,
		},
		"debug": {
			input:  "debug",
			expect: slog.LevelDebug,
		},
		"DEBUG": {
			input:  "DEBUG",
			expect: slog.LevelDebug,
		},
		"info": {
			input:  "info",
			expect: slog.LevelInfo,
		},
		"INFO": {
			input:  "INFO",
			expect: slog.LevelInfo,
		},
		"warn": {
			input:  "warn",
			expect: slog.LevelWarn,
		},
		"WARN": {
			input:  "WARN",
			expect: slog.LevelWarn,
		},
		"error": {
			input:  "error",
			expect: slog.LevelError,
		},
		"ERROR": {
			input:  "ERROR",
			expect: slog.LevelError,
		},
		"invalid": {
			input:            "INVALID_LEVEL",
			expect:           slog.LevelInfo,
			errorIs:          ErrInvalidLogLevel,
			errorMsgContains: `"INVALID_LEVEL" is not a valid log level`,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := ParseLogLevel(test.input)

			if got != test.expect {
				t.Errorf("expected log level to be %v but got %v", test.expect, got)
			}

			if test.errorIs != nil || test.errorMsgContains != "" {
				if err == nil {
					t.Fatalf("expected an error but no error")
				}

				if test.errorIs != nil && !errors.Is(err, test.errorIs) {
					t.Errorf("expected error `%v` to be `%v`", err, test.errorIs)
				}

				if test.errorMsgContains != "" && !strings.Contains(err.Error(), test.errorMsgContains) {
					t.Errorf("expected error message %q to contain %q", err.Error(), test.errorMsgContains)
				}

			} else {
				if err != nil {
					t.Fatalf("expected no error but got error: %v", err)
				}
			}
		})
	}
}

func TestParseLogLevelWithDefault(t *testing.T) {
	leveldebug, err := ParseLogLevelWithDefault("DEBUG", slog.LevelWarn)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if leveldebug != slog.LevelDebug {
		t.Errorf("expected debug log level but got %v", leveldebug)
	}

	levelwarn, err := ParseLogLevelWithDefault("", slog.LevelWarn)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if levelwarn != slog.LevelWarn {
		t.Errorf("expected warn log level but got %v", levelwarn)
	}

	levelerror, err := ParseLogLevelWithDefault("INVALID_LEVEL", slog.LevelError)
	if err == nil {
		t.Errorf("expected ErrInvalidLogLevel but did not get any error")
	}
	if !errors.Is(err, ErrInvalidLogLevel) {
		t.Errorf("expected ErrInvalidLogLevel but got %v", err)
	}
	if levelerror != slog.LevelError {
		t.Errorf("expected warn log level but got %v", levelwarn)
	}
}
