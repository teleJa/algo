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

// 567 s2是否包含s1的排列
func checkInclusion(s1 string, s2 string) bool {

	need := make(map[rune]int)
	window := make(map[rune]int)

	for _, s := range s1 {
		need[s] += 1
	}

	valid := 0
	left, right := 0, 0
	for right < len(s2) {
		r := rune(s2[right])
		right++

		if _, exist := need[r]; exist {
			window[r]++
			if window[r] == need[r] {
				valid++
			}
		}

		// 缩小窗口(保证窗口内的字符都是s1的)
		for right-left >= len(s1) {

			if valid == len(need) {
				return true
			}

			l := rune(s2[left])
			left++
			if _, exist := need[l]; exist {
				if window[l] == need[l] {
					valid--
				}
				window[l]--
			}
		}

	}
	return false
}

// 438
func findAnagrams(s string, p string) []int {
	need := make(map[rune]int)
	window := make(map[rune]int)
	for _, p0 := range p {
		need[p0]++
	}

	left, right := 0, 0
	valid := 0
	var res []int
	for right < len(s) {
		r := rune(s[right])
		right++
		if _, exist := need[r]; exist {
			window[r]++
			if window[r] == need[r] {
				valid++
			}
		}

		for right-left >= len(p) {

			if valid == len(need) {
				res = append(res, left)
			}

			l := rune(s[left])
			left++
			if _, exist := window[l]; exist {
				if window[l] == need[l] {
					valid--
				}
				window[l]--
			}

		}

	}
	return res
}
