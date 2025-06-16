package app

import (
	"time"

	"avui/internal/config"
	"avui/internal/fetcher"
	"avui/internal/generator"
	"avui/internal/models"
	"avui/internal/transformers"

	"github.com/rs/zerolog"
)

type App struct {
	Logger    zerolog.Logger
	Config    *config.Config
	Fetcher   fetcher.Fetcher
	Generator generator.Generator
}

func NewApp(logger zerolog.Logger, cfg *config.Config) (*App, error) {
	client := fetcher.DefaultHTTPClient{}
	parser := &fetcher.RSSParser{}
	feedsTransforms := []transformers.Transformer{
		&transformers.FilterRecentNews{MaxAge: cfg.MaxAge},
		&transformers.LimitItems{MaxItems: cfg.MaxItems},
	}

	rssFetcher := fetcher.New(logger, client, parser, feedsTransforms)

	tmpl, err := generator.ParseTemplate(cfg.TemplatePath)
	if err != nil {
		return nil, err
	}
	fs := generator.OSFileSystem{}
	gntr := generator.New(logger, tmpl, fs)

	return &App{
		Logger:    logger,
		Config:    cfg,
		Fetcher:   rssFetcher,
		Generator: gntr,
	}, nil
}

func (a *App) Run() error {
	a.Logger.Info().Msg("fetching feeds...")
	feeds, err := fetcher.FetchAll(a.Fetcher, a.Config, a.Logger)
	if err != nil {
		return err
	}

	page := &models.Page{
		Title:       "notícies",
		LastUpdated: time.Now().In(a.Config.Location),
		Feeds:       feeds,
	}

	a.Logger.Info().Msg("rendering file...")
	if err := a.Generator.RenderPage(page, a.Config.OutputDir); err != nil {
		return err
	}

	return nil
}
