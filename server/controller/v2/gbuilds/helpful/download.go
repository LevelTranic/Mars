package helpful

import "github.com/grafana/regexp"

func IsNumeric(s string) bool {
	re := regexp.MustCompile(`^\d+$`)
	return re.MatchString(s)
}
