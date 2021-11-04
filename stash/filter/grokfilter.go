package filter

import "regexp"

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
		reg := regexp.Compile(match_str)
		reg.search(s)
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
