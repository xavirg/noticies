package generator

import (
	"avui/internal/models"
)

type Generator interface {
	RenderPage(page *models.Page, outputPath string) error
}
