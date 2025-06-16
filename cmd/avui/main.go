package main

import (
	"flag"

	"avui/internal/app"
	"avui/internal/bootstrap"
)

func main() {
	logger := bootstrap.InitLogger()

	cfgPath := flag.String("config", "config.yaml", "Path to config file")
	flag.Parse()

	logger.Info().Str("config_path", *cfgPath).Msg("loading configuration...")
	cfg, err := bootstrap.LoadConfig(*cfgPath)
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to load configuration")
	}

	a, err := app.NewApp(logger, cfg)
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to initialize application")
	}

	if err = a.Run(); err != nil {
		logger.Fatal().Err(err).Msg("application error")
	}

	logger.Info().Str("output_dir", cfg.OutputDir).Msg("file successfully rendered")
}
