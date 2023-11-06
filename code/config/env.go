package config

import (
	"fmt"
	"log/slog"
	"os"
)

func ValidateENV() error {
	database_file := os.Getenv("DATABASE_FILE")
	slog.Info("USING: " + database_file)
	if database_file == "" {
		return fmt.Errorf("need environment variable DATABASE_FILE")
	}

	tmdb_key := os.Getenv("TMDB_KEY")
	slog.Info("TMDB_KEY: " + tmdb_key)
	if tmdb_key == "" {
		return fmt.Errorf("need environment variable TMDB_KEY")
	}

	session_key := os.Getenv("SESSION_KEY")
	slog.Info("SESSION KEY: " + session_key)
	if session_key == "" {
		return fmt.Errorf("need environment variable SESSION_KEY")
	}

	return nil
}
