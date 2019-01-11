package palindrome2integer

func isPalindrome(x int) bool {
	if x < 0 || (x != 0 && x%10 == 0) {
		return false
	}
	var sum int = 0
	for x > sum {
		sum = sum*10 + x%10
		x /= 10
	}
	return (x == sum) || (x == sum/10)
}
