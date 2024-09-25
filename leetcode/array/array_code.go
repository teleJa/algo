package array

import (
	"math"
	"regexp"
	"sort"
	"strings"
)

// 但凡数组有序都可以尝试双指针
// 此处使用快慢指针主要是为了原地修改数组

// 26 删除有序数组中的重复项
// 有序的数组,重复的元素必然是连着的
func removeDuplicates(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
		fast++
	}
	return slow + 1
}

// 27 移除所有等于val的元素
func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return len(nums)
	}

	slow, fast := -1, 0
	for fast < len(nums) {
		if nums[fast] != val {
			slow++
			nums[slow] = nums[fast]
		}
		fast++
	}
	return slow + 1
}

// 283 移动零,将所有等于0的元素移动到末尾
func moveZeroes(nums []int) {
	if len(nums) == 0 {
		return
	}
	slow, fast := -1, 0
	for fast < len(nums) {
		if nums[fast] != 0 {
			slow++
			nums[slow] = nums[fast]
		}
		fast++
	}

	for i := slow + 1; i < len(nums); i++ {
		nums[i] = 0
	}

}

// 此处双向指针做二分查找
// 167(注意题目要求数组下标从1开始)
func twoSum(numbers []int, target int) []int {

	left, right := 0, len(numbers)-1

	for left < right {
		if numbers[left]+numbers[right] == target {
			return []int{left + 1, right + 1}
		}
		if numbers[left]+numbers[right] > target {
			right--
		}
		if numbers[left]+numbers[right] < target {
			left++
		}
	}
	return nil
}

// 双指针交换数据
// 344 反转字符串
func reverseString(s []byte) {

	left, right := 0, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}

// 125 判断是否为回文串
func isPalindrome(s string) bool {

	// 去除非数字和字母的部分
	compile, _ := regexp.Compile("[^a-zA-Z0-9]+")
	s = compile.ReplaceAllString(s, "")

	left, right := 0, len(s)-1
	for left < right {

		if !strings.EqualFold(string(s[left]), string(s[right])) {
			return false
		}
		left++
		right--
	}
	return true

}

// 5 最长回文子串 (耗时较长,无法ac)
func longestPalindrome2(s string) string {
	if len(s) < 2 {
		return s
	}

	var res string
	left, right := 0, 1
	for left < len(s) {
		for right <= len(s) {
			sub := s[left:right]
			if isPalindrome(sub) && len(sub) > len(res) {
				res = sub
			}
			right++
		}
		left++
		right = left
	}
	return res
}

// 5 最长回文子串
func longestPalindrome(s string) string {

	var res string
	for i := 0; i < len(s); i++ {
		odd := findPalindrome(s, i, i)
		even := findPalindrome(s, i, i+1)

		if len(odd) > len(res) {
			res = odd
		}

		if len(even) > len(res) {
			res = even
		}

	}
	return res
}

// 从中心向两边的双指针
func findPalindrome(s string, left, right int) string {

	for left >= 0 && right < len(s) && strings.EqualFold(string(s[left]), string(s[right])) {
		left--
		right++
	}
	return s[left+1 : right]
}

// 88. 合并两个有序数组
func merge(nums1 []int, m int, nums2 []int, n int) {

	// 从后向前合并(nums1数组后面是空的)
	m0, n0 := m-1, n-1
	p := m + n - 1
	for m0 >= 0 && n0 >= 0 {
		if nums1[m0] < nums2[n0] {
			nums1[p] = nums2[n0]
			n0--
		} else {
			nums1[p] = nums1[m0]
			m0--
		}
		p--
	}

	// 判断num2是否遍历完成
	for n0 >= 0 {
		nums1[p] = nums2[n0]
		n0--
		p--
	}
}

// 977. 有序数组的平方
// nums中可能存在负数,平方后变成正数
func sortedSquares(nums []int) []int {

	// m0 代表左侧小于0数据 n0代表右侧大于0数据
	m0, n0 := 0, 0
	for i := range nums {
		if nums[i] >= 0 {
			n0 = i
			break
		} else {
			n0++
		}
	}

	m0 = n0 - 1
	res := make([]int, len(nums))
	index := 0
	// 合并数组
	for m0 > -1 && n0 < len(nums) {
		abs := math.Abs(float64(nums[m0]))
		if abs < float64(nums[n0]) {
			res[index] = int(abs)
			m0--
		} else {
			res[index] = nums[n0]
			n0++
		}
		index++
	}

	for m0 > -1 {
		abs := math.Abs(float64(nums[m0]))
		res[index] = int(abs)
		m0--
		index++
	}

	for n0 < len(nums) {
		res[index] = nums[n0]
		n0++
		index++
	}

	for i := range res {
		res[i] = res[i] * res[i]
	}
	return res

}

// 双指针初始化在数组的尾部,从绝对值来看,左侧负数子数组从左向右是递减,右侧正数子数组从右向左也是递减(v字形)
func sortedSquares2(nums []int) []int {

	// 此时left,right 代表2个子数组绝对值最大的索引
	left, right := 0, len(nums)-1
	res := make([]int, len(nums))
	p := len(nums) - 1
	for left <= right {
		l := math.Abs(float64(nums[left]))
		r := math.Abs(float64(nums[right]))
		if l > r {
			res[p] = int(l) * int(l)
			left++
		} else {
			res[p] = int(r) * int(r)
			right--
		}
		p--
	}
	return res
}

// 360 有序转化数组
// 给你一个已经 排好序 的整数数组 nums 和整数 a 、 b 、 c 。对于数组中的每一个元素 nums[i] ，计算函数值 f(x) = ax2 + bx + c ，请 按升序返回数组 。
// 此函数为抛物线,a > 0 开口向上, a < 0 开口向下
// a=0的时候考虑 b的情况 b > 0,f(x)逐渐递增(退化为开口向下抛物线的左侧), b <0 ,f(x)逐渐递减(退化为开口向下抛物线的右侧),因此 a=0应归为抛物线开口向下的情况
func sortTransformedArray(nums []int, a, b, c int) []int {

	left, right := 0, len(nums)-1
	p := 0
	if a > 0 {
		p = right
	}

	fn := func(x, a, b, c int) int {
		return a*x*x + b*x + c
	}

	res := make([]int, len(nums))
	for left <= right {

		l0 := fn(nums[left], a, b, c)
		r0 := fn(nums[right], a, b, c)
		// 开口向上,两边向中心递减
		if a > 0 {
			if l0 > r0 {
				res[p] = l0
				p--
				left++
			} else {
				res[p] = r0
				p--
				right--
			}
		} else {
			if l0 > r0 {
				res[p] = r0
				right--
				p++
			} else {
				res[p] = l0
				left++
				p++
			}
		}

	}
	return res
}

// 1329. 将矩阵按对角线排序
// 同一对角线上的元素,横纵坐标的差相同
func diagonalSort(mat [][]int) [][]int {

	m := make(map[int][]int)
	for i := range mat {
		arr := mat[i]
		for j := range arr {
			c := j - i
			m[c] = append(m[c], mat[i][j])
		}
	}

	// 排序
	for _, v := range m {
		sort.Ints(v)
	}

	for i := range mat {
		arr := mat[i]
		for j := range arr {
			c := j - i
			mat[i][j] = m[c][0]
			m[c] = append(m[c][1:])
		}
	}
	return mat
}

// 1260 二维网格迁移
func shiftGrid(grid [][]int, k int) [][]int {

	// 映射为一维数组
	m := len(grid)
	n := len(grid[0])

	arr := make([]int, m*n)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// 一维数组下标
			index := (i*n + j + k) % len(arr)
			arr[index] = grid[i][j]
		}
	}

	// 转换为二维
	for i := range arr {
		grid[i/n][i%n] = arr[i]

	}

	return grid

}
