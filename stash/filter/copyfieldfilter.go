package filter

func CopyFiledFilter(field, target string) FilterFunc {
	return func(m map[string]interface{}) map[string]interface{} {
		val, ok := m[field]
		if !ok {
			return m
		}

		s, ok := val.(string)
		if !ok {
			return m
		}

		if len(target) > 0 {
			m[target] = s
		}

		return m
	}
}
