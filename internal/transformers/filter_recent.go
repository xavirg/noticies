package transformers

import (
	"time"

	"avui/internal/models"
)

type FilterRecentNews struct {
	MaxAge time.Duration
}

func (f *FilterRecentNews) Transform(items []*models.Item) []*models.Item {
	var filtered []*models.Item
	now := time.Now()

	for _, item := range items {
		if item.PublishedAt.IsZero() || now.Sub(item.PublishedAt) <= f.MaxAge {
			filtered = append(filtered, item)
		}
	}

	return filtered
}
