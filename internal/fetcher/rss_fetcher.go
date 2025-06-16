package fetcher

import (
	"fmt"
	"io"
	"time"

	"avui/internal/config"
	"avui/internal/models"
	"avui/internal/transformers"

	"github.com/rs/zerolog"
)

type RSSFetcher struct {
	logger       zerolog.Logger
	client       HTTPClient
	parser       Parser
	transformers []transformers.Transformer
}

func New(logger zerolog.Logger, client HTTPClient, parser Parser, transformers []transformers.Transformer) *RSSFetcher {
	return &RSSFetcher{
		logger:       logger,
		client:       client,
		parser:       parser,
		transformers: transformers,
	}
}

func (f *RSSFetcher) Fetch(feedCfg *config.Feed, location *time.Location) (*models.Feed, error) {
	resp, err := f.client.Get(feedCfg.URL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch URL: %w", err)
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			f.logger.Fatal().Err(err).Msg("could not close response body reader")
		}
	}(resp.Body)

	feed, err := f.parser.Parse(resp.Body, feedCfg, location)
	if err != nil {
		return nil, fmt.Errorf("failed to parse feed: %w", err)
	}

	for _, t := range f.transformers {
		feed.Items = t.Transform(feed.Items)
	}
	return feed, nil
}
