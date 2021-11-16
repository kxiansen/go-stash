package filter

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
)

func MutateFilter(Add_fields [][]string) FilterFunc {
	return func(m map[string]interface{}) map[string]interface{} {
		//将%{}%中的字段取出
		re1 := regexp.MustCompile(`^.*?%{\s*?(?P<d_field>\S+)\s*?}%.*`)
		//将%{}%替换成map中对应字段的值
		re2 := regexp.MustCompile(`%{\s*?\S+\s*?}%?`)
		for _, field := range Add_fields {
			field_name := field[0]
			value := field[1]
			if strings.Contains(value, "%{") {
				for {
					match := re1.FindStringSubmatch(value)
					if len(match) >= 2 {
						match = re1.FindStringSubmatch(value)
						found := re2.FindString(value)
						if found != "" {
							switch v := m[match[1]].(type) {
							case map[string]interface{}:
								vjson, _ := json.Marshal(v)
								value = strings.Replace(value, found, string(vjson), 1)
							default:
								defer func() {
									if err := recover(); err != nil {
										fmt.Println("--------------------------------------------------------------------------------------")
										fmt.Println("没有匹配到source_str: ", v)
										fmt.Printf("\n")
										fmt.Printf("match: \"%v\",len(match): %d, found: \"%v\"\n", match, len(match), found)
										fmt.Println("source data: ", m)
										fmt.Println("[error]: ", err) //这里的err其实就是panic传入的内容，"bug"
										fmt.Println("--------------------------------------------------------------------------------------")
										panic("stop...")
									}
								}()
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
