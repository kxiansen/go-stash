package filter

import (
	"regexp"
	"strings"
)

func MutateFilter(Add_fields [][]string) FilterFunc {
	return func(m map[string]interface{}) map[string]interface{} {
		for _, field := range Add_fields {
			field_name := field[0]
			value := field[1]
			if strings.Contains(value, "%{") {
				for {
					re := regexp.MustCompile(`.*%{\s*(?P<d_field>\S+)\s*}%.*`)
					match := re.FindStringSubmatch(value)
					// groupNames := re.SubexpNames()
					if len(match) >= 2 {
						// v := m[match[1]].(string) + match[2] + m[match[3]].(string)
						value = re.ReplaceAllString(value, m[match[1]].(string))
					} else {
						break
					}
				}
			}
			m[field_name] = value
		}
		return m
	}
}
