package fetcher

import (
	"encoding/xml"
	"io"
	"sort"
	"time"

	"avui/internal/config"
	"avui/internal/models"

	"golang.org/x/net/html/charset"
)

type Parser interface {
	Parse(io.Reader, *config.Feed, *time.Location) (*models.Feed, error)
}

type RSSParser struct{}

func (p *RSSParser) Parse(reader io.Reader, feedCfg *config.Feed, location *time.Location) (*models.Feed, error) {
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReaderLabel

	var rssFeed models.RssFeed
	if err := decoder.Decode(&rssFeed); err != nil {
		return nil, err
	}

	var newItems []*models.Item
	for _, item := range rssFeed.Channel.Items {
		parsedTime, err := time.Parse(feedCfg.TimeFormat, item.PubDate)
		if err != nil {
			continue
		}
		newItems = append(newItems, &models.Item{
			Title:       item.Title,
			Link:        item.Link,
			PublishedAt: parsedTime.In(location),
		})
	}

	sort.Slice(newItems, func(i, j int) bool {
		return newItems[i].PublishedAt.After(newItems[j].PublishedAt)
	})

	if rssFeed.Channel.Title == "Portada" {
		rssFeed.Channel.Title = feedCfg.Name
	}

	return &models.Feed{
		Title: rssFeed.Channel.Title,
		Items: newItems,
		Order: feedCfg.Order,
	}, nil
}
