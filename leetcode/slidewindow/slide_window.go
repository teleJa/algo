package slidewindow

// 76 最小覆盖子串
func minWindow(s string, t string) string {

	need := map[rune]int{}
	window := map[rune]int{}
	for _, t0 := range t {
		need[t0] += 1
	}

	var res string

	left, right := 0, 0
	for right < len(s) {
		if _, ok := need[rune(s[right])]; ok {
			window[rune(s[right])] += 1
		}
		// 判断此时窗口内的元素是否已经满足t
		match := matchWindow(need, window)
		// 缩小窗口
		for match {

			sub := s[left : right+1]
			if res == "" {
				res = sub
			} else if len(sub) < len(res) {
				res = sub
			}

			c := s[left]
			// 缩小窗口
			left++
			if _, exist := window[rune(c)]; exist {
				window[rune(c)]--
			}
			// 比较是否仍满足
			match = matchWindow(need, window)

		}
		right++
	}

	return res
}

func matchWindow(need, window map[rune]int) bool {
	if len(window) < len(need) {
		return false
	}
	for k, v := range need {
		if v > window[k] {
			return false
		}
	}
	return true
}
