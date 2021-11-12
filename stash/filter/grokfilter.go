package filter

import (
	"encoding/json"
	"fmt"
	"regexp"
)

func GrokFilter(field, match_str string) FilterFunc {
	return func(m map[string]interface{}) map[string]interface{} {
		val, ok := m[field]
		if !ok {
			return m
		}

		var s string
		m_tmp, ok := val.(map[string]interface{})
		if ok {
			dataType, _ := json.Marshal(m_tmp)
			s = string(dataType)

		} else {
			s, ok = val.(string)
			if !ok {
				return m
			}
		}
		fmt.Println(s)

		re := regexp.MustCompile(match_str)
		match := re.FindStringSubmatch(s)
		groupNames := re.SubexpNames()

		defer func() {
			if err := recover(); err != nil {
				fmt.Println("--------------------------------------------------------------------------------------")
				fmt.Println("source_str: ")
				fmt.Println(s)
				fmt.Printf("\n")
				fmt.Printf("match: \"%v\",groupNames: \"%v\",len(match): %d,len(groupNames): %d\n", match, groupNames, len(match), len(groupNames))
				fmt.Println("[error]: ", err) //这里的err其实就是panic传入的内容，"bug"
				fmt.Println("--------------------------------------------------------------------------------------")
				panic("stop...")
			}
		}()
		for i, name := range groupNames {
			if i != 0 && name != "" {
				m[name] = match[i]
			}
		}

		return m
	}
}
