package zigzag

func convert(s string, numRows int) string {
	if s == "" || len(s) == 0 || numRows <= 0 {
		return ""
	}
	if numRows == 1 {
		return s
	}

	var rune_s = []rune(s)
	var str_len = len(rune_s)
	var rune_res = make([]rune, str_len)
	var res_id, g_times, count = 0, 0, 0
	var gap = numRows + (numRows - 2)

	for row := 0; row < numRows; row++ {
		for {
			if row == 0 {
				var offset = row + (g_times * gap)
				if offset >= str_len {
					g_times = 0
					break
				}
				rune_res[res_id] = rune_s[offset]
				res_id += 1
				g_times += 1
			} else if row == numRows-1 {
				var offset = row + (g_times * gap)
				if offset >= str_len {
					break
				}
				rune_res[res_id] = rune_s[offset]
				res_id += 1
				g_times += 1
			} else {
				var offset int
				if count%2 == 0 {
					offset = row + (g_times * gap)
					count += 1
					g_times += 1
				} else {
					offset = row + (g_times * gap) - (2 * row)
					count += 1
				}
				if offset >= str_len {
					g_times = 0
					count = 0
					break
				}
				rune_res[res_id] = rune_s[offset]
				res_id += 1
			}
		}
	}
	return string(rune_res)
}
