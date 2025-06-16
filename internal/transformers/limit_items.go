package transformers

import "avui/internal/models"

type LimitItems struct {
	MaxItems int
}

func (l *LimitItems) Transform(items []*models.Item) []*models.Item {
	if len(items) <= l.MaxItems {
		return items
	}
	return items[:l.MaxItems]
}
