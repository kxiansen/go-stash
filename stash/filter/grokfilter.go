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
		groupNames := re.SubexpNames()

		defer func() {
			fmt.Println("--------------------------------------------------------------------------------------")

			fmt.Println(s)
			fmt.Printf("match: \"%v\",groupNames: \"%v\",len(match): %d,len(groupNames): %d\n", match, groupNames, len(match), len(groupNames))
			if err := recover(); err != nil {
				fmt.Println("[error]: ", err) //这里的err其实就是panic传入的内容，"bug"
			}
			fmt.Println("--------------------------------------------------------------------------------------")
		}()
		for i, name := range groupNames {
			if i != 0 && name != "" {
				m[name] = match[i]
			}
		}

		return m
	}
}
