package config

import (
	"fmt"
	"log/slog"
	"os"
)

func ValidateENV() error {
	database_file := os.Getenv("DATABASE_FILE")
	slog.Debug("USING: " + database_file)
	if database_file == "" {
		return fmt.Errorf("need environment variable DATABASE_FILE")
	}

	host := os.Getenv("HOST")
	slog.Debug("USING: " + host)
	if host == "" {
		return fmt.Errorf("need environment variable HOST")
	}

	tmdb_key := os.Getenv("TMDB_KEY")
	slog.Debug("TMDB_KEY: " + tmdb_key)
	if tmdb_key == "" {
		return fmt.Errorf("need environment variable HOST")
	}

	return nil
}
