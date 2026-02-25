package projector

// toIntSlice converts an any value to a slice of ints.
func toIntSlice(val any) ([]int, bool) {
	switch v := val.(type) {
	case []int:
		return v, true
	case []any:
		result := make([]int, 0, len(v))
		for _, elem := range v {
			switch e := elem.(type) {
			case float64:
				result = append(result, int(e))
			case int:
				result = append(result, e)
			default:
				return nil, false
			}
		}
		return result, true
	default:
		return nil, false
	}
}

// toInt converts an any value to int.
func toInt(val any) int {
	switch v := val.(type) {
	case float64:
		return int(v)
	case int:
		return v
	default:
		return 0
	}
}
