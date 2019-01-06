package actions

import "strings"

type Config struct {
	Extension  string
	Pattern    string
	AddTop     bool
	AddBottom  bool
	PrefixLine string
}

var defaultPattern = map[string]Config{
	".go": Config{
		Extension:  ".go",
		Pattern:    "/*\n{{content}}\n*/",
		AddTop:     true,
		AddBottom:  true,
		PrefixLine: " * ",
	},
	".js": Config{
		Extension:  ".js",
		Pattern:    "/*\n{{content}}\n*/",
		AddTop:     true,
		AddBottom:  true,
		PrefixLine: " * ",
	},
	".css": Config{
		Extension:  ".css",
		Pattern:    "/*\n{{content}}\n*/",
		AddTop:     true,
		AddBottom:  true,
		PrefixLine: " * ",
	},
	".html": Config{
		Extension:  ".html",
		Pattern:    "<!--\n{{content}}\n-->",
		AddTop:     true,
		AddBottom:  true,
		PrefixLine: " ",
	},
}

// TODO: Load from configuration
// GetCommentPatterns loads patterns for formatting the text before pushing into source code files.
func GetCommentPatterns() map[string]Config {
	return defaultPattern
}

func (c Config) GetReplacement(replacement string) string {
	result := strings.Replace(c.PrefixLine+replacement, "\n", "\n"+c.PrefixLine, -1)
	result = strings.Replace(c.Pattern, "{{content}}", result, -1)
	return result
}
