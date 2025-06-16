package generator

import (
	"html/template"
	"path/filepath"

	"avui/internal/utils"
)

func ParseTemplate(templatePath string) (*template.Template, error) {
	return template.New(filepath.Base(templatePath)).
		Funcs(utils.TemplateFuncMap()).
		ParseFiles(templatePath)
}
