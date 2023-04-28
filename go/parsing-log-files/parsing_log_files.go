package parsinglogfiles

import "regexp"

var (
	validBegining = []*regexp.Regexp{
		regexp.MustCompile("^\\[TRC\\]"),
		regexp.MustCompile("^\\[DBG\\]"),
		regexp.MustCompile("^\\[INF\\]"),
		regexp.MustCompile("^\\[WRN\\]"),
		regexp.MustCompile("^\\[ERR\\]"),
		regexp.MustCompile("^\\[FTL\\]"),
	}
)

func IsValidLine(text string) bool {
	for _, regex := range validBegining {
		if regex.MatchString(text) {
			return true
		}
	}

	return false
}

func SplitLogLine(text string) []string {
	regex := regexp.MustCompile("<[~\\*=-]*>")

	return regex.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	regex := regexp.MustCompile(`(?i)".*password.*"`)
	var count int
	for _, line := range lines {
		if regex.MatchString(line) {
			count += 1
		}
	}

	return count
}

func RemoveEndOfLineText(text string) string {
	regex := regexp.MustCompile("end-of-line\\d*")

	return regex.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	regex := regexp.MustCompile("User\\s+([a-zA-z]+\\d*)\\s+")

	result := []string{}
	for _, line := range lines {
		matches := regex.FindStringSubmatch(line)
		if matches != nil && matches[1] != "" {
			result = append(result, "[USR] "+matches[1]+" "+line)
		} else {
			result = append(result, line)
		}
	}

	return result
}
