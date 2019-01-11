package roman2Int

func romanToInt(s string) int {
	var result int
	for i := len(s) - 1; i >= 0; i-- {
		switch string(s[i]) {
		case "I":
			if result >= 5 {
				result += -1
			} else {
				result += 1
			}
		case "V":
			result += 5
		case "X":
			if result >= 50 {
				result += -10
			} else {
				result += 10
			}
		case "L":
			result += 50
		case "C":
			if result >= 500 {
				result += -100
			} else {
				result += 100
			}
		case "D":
			result += 500
		case "M":
			result += 1000
		}
	}
	return result
}
