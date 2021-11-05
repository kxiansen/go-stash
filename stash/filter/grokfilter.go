package filter

import (
	"regexp"
)

func GrokFilter(field, match_str string) FilterFunc {
	return func(m map[string]interface{}) map[string]interface{} {
		val, ok := m[field]
		if !ok {
			return m
		}

		s, ok := val.(string)
		if !ok {
			return m
		}
		re := regexp.MustCompile(match_str)
		match := re.FindStringSubmatch(s)
		groupNames := re.SubexpNames()
		// fmt.Printf("%v,%v,%d,%d\n", match, groupNames, len(match), len(groupNames))
		for i, name := range groupNames {
			if i != 0 {
				m[name] = match[i]
			}
		}

		return m
	}
}
