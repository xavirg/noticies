package bootstrap

import (
	"os"
	"time"

	"avui/internal/config"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func InitLogger() zerolog.Logger {
	return log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339})
}

func LoadConfig(path string) (*config.Config, error) {
	cfg, err := config.LoadConfig(path)
	if err != nil {
		return nil, err
	}
	loc, err := time.LoadLocation(cfg.Timezone)
	if err != nil {
		return nil, err
	}
	cfg.Location = loc
	return cfg, nil
}
