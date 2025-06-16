package fetcher

import (
	"errors"
	"sort"
	"sync"
	"time"

	"avui/internal/config"
	"avui/internal/models"

	"github.com/rs/zerolog"
)

type Fetcher interface {
	Fetch(feedCfg *config.Feed, location *time.Location) (*models.Feed, error)
}

func FetchAll(fetcher Fetcher, cfg *config.Config, logger zerolog.Logger) ([]*models.Feed, error) {
	if len(cfg.Feeds) == 0 {
		return nil, errors.New("no feeds to process")
	}

	var (
		wg    sync.WaitGroup
		mu    sync.Mutex
		feeds []*models.Feed
	)

	for _, feedCfg := range cfg.Feeds {
		wg.Add(1)
		go func(feed config.Feed) {
			defer wg.Done()

			logger.Info().Str("feed", feed.Name).Msg("fetching feed")
			feedContent, err := fetcher.Fetch(&feed, cfg.Location)
			if err != nil {
				logger.Error().Err(err).Str("url", feed.URL).Msg("failed to fetch feed")
				return
			}

			mu.Lock()
			feeds = append(feeds, feedContent)
			mu.Unlock()
		}(feedCfg)
	}

	wg.Wait()

	sort.Slice(feeds, func(i, j int) bool {
		return feeds[i].Order < feeds[j].Order
	})
	return feeds, nil
}
