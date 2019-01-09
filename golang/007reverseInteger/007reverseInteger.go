package reverseInterger

func reverse(x int) int{
	const (
		MaxInt = 1 << 31 -1
		MinInt = -MaxInt - 1
	)
	var result int64
	var tmp int64 = int64(x)
	for tmp != 0{
		result = result*10 +tmp%10
		tmp /= 10
	}
	if !( result>int64(MaxInt) || result<int64(MinInt) ) { return int(result) }
	return 0

}
