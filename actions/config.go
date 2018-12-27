package actions

import "strings"

type Config struct {
	Extension string
	Pattern   string
	AddTop    bool
	AddBottom bool
}

var defaultPattern = map[string]Config{
	".go": Config{
		Extension: ".go",
		Pattern:   "/*\n{{content}}\n*/",
		AddTop:    true,
		AddBottom: true,
	},
	".js": Config{
		Extension: ".js",
		Pattern:   "/*\n{{content}}\n*/",
		AddTop:    true,
		AddBottom: true,
	},
	".css": Config{
		Extension: ".css",
		Pattern:   "/*\n{{content}}\n*/",
		AddTop:    true,
		AddBottom: true,
	},
	".html": Config{
		Extension: ".html",
		Pattern:   "<!--\n{{content}}\n-->",
		AddTop:    true,
		AddBottom: true,
	},
}

// TODO: Load from configuration
// GetCommentPatterns loads patterns for formatting the text before pushing into source code files.
func GetCommentPatterns() map[string]Config {
	return defaultPattern
}

func (c Config) GetReplacement(replacement string) string {
	return strings.Replace(c.Pattern, "{{content}}", replacement)
}
