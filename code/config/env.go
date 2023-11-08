package config

import (
	"fmt"
	"log/slog"
	"os"
	"strconv"
)

func ValidateENV() error {
	database_file := os.Getenv(EnvDatabasFile)
	slog.Info("USING: " + database_file)
	if database_file == "" {
		return fmt.Errorf("need environment variable %s", EnvDatabasFile)
	}

	tmdb_key := os.Getenv(EnvTMDBKey)
	slog.Info("TMDB_KEY: " + tmdb_key)
	if tmdb_key == "" {
		return fmt.Errorf("need environment variable %s", EnvTMDBKey)
	}

	session_key := os.Getenv(EnvSessionKey)
	slog.Info("SESSION KEY: " + session_key)
	if session_key == "" {
		return fmt.Errorf("need environment variable %s", EnvSessionKey)
	}


	enable_analytics := os.Getenv(EnvEnableAnalytics)
	enable_analytics_bool, _ := strconv.ParseBool(enable_analytics)
	if enable_analytics_bool {
		plaus_url := os.Getenv(EnvPlausibleURL)
		if plaus_url == "" {
			return fmt.Errorf("need environment variable %s if analytics is enabled", EnvPlausibleURL)
		}

		app_domain := os.Getenv(EnvAppDomain)
		if app_domain == "" {
			return fmt.Errorf("need environment variable %s if analytics is enabled", EnvAppDomain)
		}

		app_url := os.Getenv(EnvAppURL)
		if app_url == "" {
			return fmt.Errorf("need environment variable %s if analytics is enabled", EnvAppURL)
		}
	}

	return nil
}
