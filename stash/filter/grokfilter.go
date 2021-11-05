package filter

import (
	"fmt"
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
		groupName := re.SubexpNames()
		fmt.Printf("%v,%v,%d,%d\n", match, groupName, len(match), len(groupName))

		// fmt.Println(s)

		// delete(m, field)
		// if len(target) > 0 {
		// 	m[target] = nm
		// } else {
		// 	for k, v := range nm {
		// 		m[k] = v
		// 	}
		// }

		return m
	}
}
