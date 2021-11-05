package filter

import "regexp"

func ReplaceStrFilter(gsubs [][]string) FilterFunc {
	return func(m map[string]interface{}) map[string]interface{} {
		for _, gsub := range gsubs {
			field := gsub[0]
			val, ok := m[field]
			if !ok {
				return m
			}
			s, ok := val.(string)
			if !ok {
				return m
			}
			re, _ := regexp.Compile(gsub[1])
			result := re.ReplaceAllString(s, gsub[2])
			m[field] = result

		}
		return m
	}
}
