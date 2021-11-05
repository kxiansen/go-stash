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
				re := regexp.MustCompile(`%{(?P<field1>\S+)}%(?P<str1>[\s\S]*)%{(?P<field2>\S+)}%`)
				match := re.FindStringSubmatch(value)
				if len(match) == 4 {
					v := m[match[1]].(string) + match[2] + m[match[3]].(string)
					m[field_name] = v
				} else if len(match) == 3 {
					v := m[match[1]].(string) + match[2]
					m[field_name] = v
				} else if len(match) == 2 {
					v := m[match[1]].(string)
					m[field_name] = v
				}

			} else {
				m[field_name] = value
			}

		}
		return m
	}
}
