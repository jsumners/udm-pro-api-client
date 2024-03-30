package slug

import "regexp"

var re *regexp.Regexp = regexp.MustCompile(`[^a-zA-Z0-9-]+`)

// Hostname slugifies a string into a hostname compatible string.
// For example, "foo bar" becomes "foo-bar".
func Hostname(input string) string {
	return re.ReplaceAllString(input, "-")
}
