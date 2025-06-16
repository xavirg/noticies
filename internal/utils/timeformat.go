package utils

import (
	"fmt"
	"time"
)

func TemplateFuncMap() map[string]interface{} {
	return map[string]interface{}{
		"formatTime": FormatTime,
		"humanTime":  HumanizeTime,
	}
}

func FormatTime(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.Format("2006-01-02 15:04 MST")
}

func HumanizeTime(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return humanizeTime(t)
}

func humanizeTime(t time.Time) string {
	now := time.Now()
	diff := now.Sub(t)

	switch {
	case diff < time.Minute:
		return "just now"
	case diff < time.Hour:
		return pluralize(int(diff.Minutes()), "min")
	case diff < 24*time.Hour:
		return pluralize(int(diff.Hours()), "h")
	default:
		return pluralize(int(diff.Hours()/24), "d")
	}
}

func pluralize(value int, unit string) string {
	if value == 1 {
		return fmt.Sprintf("1 %s ago", unit)
	}
	return fmt.Sprintf("%d %ss ago", value, unit)
}
