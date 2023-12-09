package config

import (
	"log/slog"
	"strings"
	"time"

	"github.com/spf13/viper"
)

type MuralConfig struct {
	TodayTheme       string
	BoardWidth       int
	MaxScore         int
	Host             string `mapstructure:"HOST"`
	Version          string `mapstructure:"VERSION"`
	Env              string `mapstructure:"ENV"`
	EnableTLS        bool   `mapstructure:"ENABLE_TLS"`
	DatabaseFile     string `mapstructure:"DATABASE_FILE"`
	MigrationsFolder string `mapstructure:"MIGRATIONS_FOLDER"`
	TMDBKey          string `mapstructure:"TMDB_KEY"`
	SessionKey       string `mapstructure:"SESSION_KEY"`
	// used for a manual reset of the game
	OverrideKey        string `mapstructure:"OVERRIDE_KEY"`
	EnabledAnalytics   bool   `mapstructure:"ENABLE_ANALYTICS"`
	PlausibleURL       string `mapstructure:"PLAUS_URL"`
	PlausibleAppDomain string `mapstructure:"APP_DOMAIN"`
	AppURL             string `mapstructure:"APP_URL"`
}

const (
	EnvTest = "test"
	EnvProd = "prod"
)

func NewMuralConfig() (MuralConfig, error) {
	config := MuralConfig{}

	viper.AddConfigPath(".")
	viper.SetConfigName("mural")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		return config, err
	}

	for _, key := range viper.AllKeys() {
		envKey := strings.ToUpper(strings.ReplaceAll(key, ".", "_"))
		err := viper.BindEnv(key, envKey)
		if err != nil {
			return config, err
		}
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	config.TodayTheme = GetTodayThemeDefault()
	config.BoardWidth = 10
	config.MaxScore = 100
	return config, nil
}

const (
	Theme2020   = "2020"
	Theme2010   = "2010"
	Theme2000   = "2000"
	Theme1990   = "1990"
	Theme1980   = "1980"
	Theme1970   = "1970"
	ThemeRandom = "Random"
)

var (
	ThemeOptions = []string{
		Theme2020,
		Theme2010,
		Theme2000,
		Theme1990,
		Theme1980,
		Theme1970,
		ThemeRandom,
	}

	DecadeOptions = []string{
		Theme2020,
		Theme2010,
		Theme2000,
		Theme1990,
		Theme1980,
		Theme1970,
	}
)

func GetTodayThemeDefault() string {
	current_day := time.Now().Weekday()
	loc, _ := time.LoadLocation("America/New_York")

	if loc != nil {
		slog.Info(loc.String())
		current_day = time.Now().In(loc).Weekday()
	}

	switch current_day {
	case time.Monday:
		return Theme2020
	case time.Tuesday:
		return Theme2010
	case time.Wednesday:
		return Theme2000
	case time.Thursday:
		return Theme1990
	case time.Friday:
		return Theme1980
	case time.Saturday:
		return Theme1970
	default:
		// Sunday or any other day
		return ThemeRandom
	}
}
