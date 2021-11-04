package filter

import (
	"fmt"

	jsoniter "github.com/json-iterator/go"
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
		fmt.Println("ssssss", s)

		var nm map[string]interface{}
		if err := jsoniter.Unmarshal([]byte(s), &nm); err != nil {
			return m
		}
		fmt.Println(nm)
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
