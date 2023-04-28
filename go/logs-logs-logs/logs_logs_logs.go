package logs

import (
	"strings"
	"unicode/utf8"
)

var (
	runeToTitle = map[rune]string{
		'❗': "recommendation",
		'🔍': "search",
		'☀': "weather",
	}
)

// Application identifies the application emitting the given log.
func Application(log string) string {
	appName := "default"

	for _, char := range log {
		title, ok := runeToTitle[char]
		if ok {
			appName = title
			break
		}
	}

	return appName
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	return strings.ReplaceAll(log, string(oldRune), string(newRune))
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return limit >= utf8.RuneCountInString(log)
}
