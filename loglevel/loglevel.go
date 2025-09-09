package loglevel

import (
	"errors"
	"fmt"
	"log/slog"
	"strings"
)

var (
	ErrInvalidLogLevel = errors.New("invalid log level")
)

func ParseLogLevel(s string) (slog.Level, error) {
	upper := strings.ToUpper(s)
	switch upper {
	case "DEBUG":
		return slog.LevelDebug, nil
	case "INFO", "":
		return slog.LevelInfo, nil
	case "WARN":
		return slog.LevelWarn, nil
	case "ERROR":
		return slog.LevelError, nil
	default:
		return slog.LevelInfo, errors.Join(ErrInvalidLogLevel, fmt.Errorf("%q is not a valid log level", s))
	}
}

func ParseLogLevelWithDefault(s string, def slog.Level) (slog.Level, error) {
	if s == "" {
		return def, nil
	}

	level, err := ParseLogLevel(s)
	if err != nil {
		return def, err
	}

	return level, nil
}

func MustParseLogLevel(s string) slog.Level {
	level, err := ParseLogLevel(s)
	if err != nil && !errors.Is(err, ErrInvalidLogLevel) {
		panic(err)
	}
	return level
}

func MustParseLogLevelWithDefault(s string, def slog.Level) slog.Level {
	level, err := ParseLogLevelWithDefault(s, def)
	if err != nil && !errors.Is(err, ErrInvalidLogLevel) {
		panic(err)
	}
	return level
}
