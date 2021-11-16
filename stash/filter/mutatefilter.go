package filter

import (
	"encoding/json"
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
					re1 := regexp.MustCompile(`.*%{\s*?(?P<d_field>\S+)\s*?}%.*`)
					re2 := regexp.MustCompile(`%{\s*?\S+\s*?}%?`)
					match := re1.FindStringSubmatch(value)
					if len(match) >= 2 {
						match = re1.FindStringSubmatch(value)
						// fmt.Println("----------------------------------------------")
						found := re2.FindString(value)
						if found != "" {
							switch v := m[match[1]].(type) {
							case map[string]interface{}:
								vjson, _ := json.Marshal(v)
								value = strings.Replace(value, found, string(vjson), 1)
							default:
								value = strings.Replace(value, found, v.(string), 1)
							}
						}
						// value = re2.ReplaceAllString(value, m[match[1]].(string))
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
