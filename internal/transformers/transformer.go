package transformers

import (
	"avui/internal/models"
)

type Transformer interface {
	Transform(items []*models.Item) []*models.Item
}
