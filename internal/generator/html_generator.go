package generator

import (
	"fmt"
	"html/template"
	"path/filepath"

	"avui/internal/models"

	"github.com/rs/zerolog"
)

type HTMLGenerator struct {
	logger zerolog.Logger
	tmpl   *template.Template
	fs     FileSystem
}

func New(logger zerolog.Logger, tmpl *template.Template, fs FileSystem) *HTMLGenerator {
	return &HTMLGenerator{
		logger: logger,
		tmpl:   tmpl,
		fs:     fs,
	}
}

func (g *HTMLGenerator) RenderPage(page *models.Page, outputPath string) error {
	dir := filepath.Dir(outputPath)
	if err := g.fs.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	file, err := g.fs.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer func() {
		if cerr := file.Close(); cerr != nil {
			g.logger.Error().Err(cerr).Msg("Failed to close output file")
		}
	}()

	if err := g.tmpl.Execute(file, page); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	return nil
}
