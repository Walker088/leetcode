package DivideTwoIntegers

import (
	"fmt"
	"testing"
)

func Test_DivideTwoIntegers(t *testing.T) {
	fmt.Println("Test_DivideTwoIntegers")
	testCases := map[int]int{7: -3, 10: 3}
	for dvd, dvs := range testCases {
		res := DivideTwoIntegers(dvd, dvs)
		println("divident: ", dvd, "divisor: ", dvs, "Res: ", res)
	}
}
