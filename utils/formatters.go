package utils

import (
	"strings"
)

// Util function for formatting aliases
func FormatAlias(alias string) string {
	formattedAlias := strings.ToLower(alias)
	formattedAlias = strings.ReplaceAll(formattedAlias, "-", "_")
	formattedAlias = strings.ReplaceAll(formattedAlias, " ", "_")
	return formattedAlias
}
