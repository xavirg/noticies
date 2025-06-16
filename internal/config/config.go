package config

import (
	"fmt"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Feeds        []Feed        `yaml:"feeds"`
	OutputDir    string        `yaml:"outputDir"`
	TemplatePath string        `yaml:"templatePath"`
	Timezone     string        `yaml:"timezone"`
	MaxAge       time.Duration `yaml:"maxAge"`
	MaxItems     int           `yaml:"maxItems"`
	Location     *time.Location
}

func (c *Config) Validate() error {
	if len(c.Feeds) == 0 {
		return fmt.Errorf("at least one feed must be specified")
	}
	if c.OutputDir == "" {
		return fmt.Errorf("outputDir is required")
	}
	if c.TemplatePath == "" {
		return fmt.Errorf("templatePath is required")
	}
	if c.Timezone == "" {
		return fmt.Errorf("timezone is required (tz database format)")
	}
	if c.MaxAge <= 0 {
		return fmt.Errorf("maxAge must be a duration (e.g. 24h, 5m)")
	}
	if c.MaxItems <= 0 {
		return fmt.Errorf("maxItems must be positive")
	}
	for i, feed := range c.Feeds {
		if err := feed.Validate(); err != nil {
			return fmt.Errorf("feed %d (%s): %w", i, feed.Name, err)
		}
	}
	return nil
}

type Feed struct {
	Name       string `yaml:"name"`
	URL        string `yaml:"url"`
	TimeFormat string `yaml:"time_format"`
	Order      int    `yaml:"order"`
}

func (f *Feed) Validate() error {
	if f.Name == "" {
		return fmt.Errorf("name is required")
	}
	if f.URL == "" {
		return fmt.Errorf("url is required")
	}
	if f.TimeFormat == "" {
		return fmt.Errorf("time_format is required")
	}
	return nil
}

func LoadConfig(filename string) (*Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	if err := config.Validate(); err != nil {
		return nil, fmt.Errorf("invalid config: %w", err)
	}

	return &config, nil
}
