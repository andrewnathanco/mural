package config

import "log/slog"

func Must(err error) {
	if err != nil {
		slog.Error(err.Error())
		panic(1)
	}
}
