package DivideTwoIntegers

import "math"

//DivideTwoIntegers returns a integer
func DivideTwoIntegers(dividend int, divisor int) int {
	/* Given two integers dividend and divisor,
	   divide two integers without using multiplication,
	   division and mod operator.
	   Return the quotient after dividing dividend by divisor.
	   The integer division should truncate toward zero. */
	// Check the special cases:
	// 1. divident equals to 2^31 - 1
	// 2. divident equals to zero
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}
	if dividend == 0 {
		return 0
	}

	// Check whether the result should be negative
	var sign = -1
	var res = 0
	if (dividend > 0 && divisor > 0) || (dividend < 0 && divisor < 0) {
		sign = 1
	}
	dvd, dvs := convertToInt64(dividend), convertToInt64(divisor)
	for dvd >= dvs {
		var shift uint
		for dvd >= (dvs << shift) {
			shift++
		}
		res += (1 << (shift - 1))
		dvd -= dvs << (shift - 1)
	}
	return res * sign
}

func convertToInt64(n int) int64 {
	res := int64(n)
	if res < 0 {
		return -res
	}
	return res
}
