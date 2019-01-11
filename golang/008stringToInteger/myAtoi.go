package myAtoi

import "fmt"

const (
	maxInt32 = (1 << 31) - 1
	minInt32 = -1 << 31
)

func IsNumber(chunks rune) bool {
	if chunks > 57 || chunks < 48 {
		return false
	} else {
		return true
	}
}

func myAtoi(str string) int {
	fmt.Println("input str:", str, "length:", len(str))
	var str_chunks = []rune(str)
	var sign, res = 0, 0
	for i := 0; i < len(str_chunks); i++ {
		if string(str_chunks[i]) == " " {
			fmt.Println("space, utf8code:", str_chunks[i])
			continue
		}
		// add sign
		if string(str_chunks[i]) == "+" && sign == 0 {
			fmt.Println("sign = 1, utf8code:", str_chunks[i])
			if (i != len(str_chunks)-1) && (IsNumber(str_chunks[i+1]) != true) {
				break
			}
			sign = 1
			continue
		} else if string(str_chunks[i]) == "-" && sign == 0 {
			fmt.Println("sign = -1, utf8code", str_chunks[i])
			if (i != len(str_chunks)-1) && (IsNumber(str_chunks[i+1]) != true) {
				break
			}
			sign = -1
			continue
		}

		// return if first character is not integer
		if IsNumber(str_chunks[i]) == false {
			fmt.Println("Return, cuz charactor is not integer:", string(str_chunks[i]))
			break
		}

		if IsNumber(str_chunks[i]) == true {
			fmt.Println("int:", string(str_chunks[i]), "utf8code:", str_chunks[i])
			fmt.Println("res:", res, "character number:", int(str_chunks[i])-48)
			fmt.Println("=========================")
			//tmp := (int(str_chunks[i]) - 48)
			res = (10 * res) + (int(str_chunks[i]) - 48)
			if (i != len(str_chunks)-1) && (IsNumber(str_chunks[i+1]) != true) {
				break
			}
			if res >= maxInt32 || -res <= minInt32 {
				fmt.Println("overflow, break")
				break
			}
		}
	}
	if sign != 1 && sign != -1 {
		sign = 1
	}
	if (sign == 1 || sign == 0) && res >= maxInt32 {
		fmt.Println("overflow, return maxInt32")
		return maxInt32
	} else if sign == -1 && -res <= minInt32 {
		fmt.Println("overflow, return minInt32")
		return minInt32
	}

	fmt.Println("res:", res, "sign:", sign)
	return (res * sign)
}
