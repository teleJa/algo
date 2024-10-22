package array

import (
	"math"
	"regexp"
	"slices"
	"sort"
	"strings"
	"unicode"
)

// 但凡数组有序都可以尝试双指针
// 此处使用快慢指针主要是为了原地修改数组

// 26 删除有序数组中的重复项
// 有序的数组,重复的元素必然是连着的
func removeDuplicates(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	// 移动元素可以理解为从0开始对原数组添加元素
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

// 80 删除有序数组中的重复项
func removeDuplicates2(nums []int) int {

	slow, fast := 0, 0
	k := 2
	// 重复次数
	count := 0
	for fast < len(nums) {
		if nums[slow] != nums[fast] {
			slow++
			nums[slow] = nums[fast]
		} else if slow < fast && count < k { // 上面if未命中,说明此时nums[slow]=nums[fast],故count需小于k,而不是小于等于
			slow++
			nums[slow] = nums[fast]
		}
		count++
		// 重置count
		if fast < len(nums)-1 && nums[fast] != nums[fast+1] {
			count = 0
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
// 167(注意题目要求数组下标从1开始,非递减数组)
func twoSum2(numbers []int, target int) []int {

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

type Val struct {
	value, index int
}

type ValArr []Val

func (v ValArr) Len() int {
	return len(v)
}

func (v ValArr) Less(i, j int) bool {
	return v[i].value < v[j].value
}

func (v ValArr) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}

// 1.两数之和
func twoSum(nums []int, target int) []int {

	var array []Val
	for i := range nums {
		array = append(array, Val{
			nums[i], i,
		})
	}
	// 排序
	sort.Sort(ValArr(array))
	left, right := 0, len(nums)-1

	for left < right {
		valLeft := array[left]
		valRight := array[right]

		if valLeft.value+valRight.value == target {
			return []int{valLeft.index, valRight.index}
		}
		if valLeft.value+valRight.value > target {
			right--
		}
		if valLeft.value+valRight.value < target {
			left++
		}
	}
	return nil
}

// 15.三数之和
func threeSum(nums []int) [][]int {
	// 三数和为0
	target := 0
	res := make([][]int, 0)

	// 排序
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		twoNumSum := target - nums[i]
		// 从i+1开始寻找两数之和
		newNums := make([]int, len(nums)-i-1)
		copy(newNums, nums[i+1:])
		left, right := 0, len(newNums)-1
		for left < right {
			sum := newNums[left] + newNums[right]
			if sum < twoNumSum {
				// 跳过重复元素
				l := newNums[left]
				for left < right && newNums[left] == l {
					left++
				}
			} else if sum > twoNumSum {
				r := newNums[right]
				for left < right && newNums[right] == r {
					right--
				}
			} else {
				res = append(res, []int{nums[i], newNums[left], newNums[right]})
				l := newNums[left]
				for left < right && newNums[left] == l {
					left++
				}
				r := newNums[right]
				for left < right && newNums[right] == r {
					right--
				}
			}
		}

		// 此种方式跳过重复元素,执行for循环时i会再次++
		/*cur := nums[i]
		for i < len(nums) && nums[i] == cur {
			i++
		}*/

		// 跳过重复元素(第一个数不重复即可保证整体不重复)
		for i < len(nums)-1 && nums[i] == nums[i+1] {
			i++
		}
	}
	return res
}

// n数之和
func nSumTarget(nums []int, n int, start int, target int64) [][]int {
	res := make([][]int, 0)
	if n < 2 || len(nums) < n {
		return res
	}

	// 两数之和
	if n == 2 {
		left, right := start, len(nums)-1
		for left < right {
			lv, rv := nums[left], nums[right]
			if int64(lv+rv) == target {
				res = append(res, []int{lv, rv})
				for left < right && nums[left] == lv {
					left++
				}
				for left < right && nums[right] == rv {
					right--
				}
			} else if int64(lv+rv) > target {
				for left < right && nums[right] == rv {
					right--
				}
			} else {
				for left < right && nums[left] == lv {
					left++
				}
			}
		}
		return res
	} else {
		// 递归
		for i := start; i < len(nums); i++ {

			sumTarget := nSumTarget(nums, n-1, i+1, target-int64(nums[i]))
			for _, r := range sumTarget {
				r = append(r, nums[i])
				res = append(res, r)
			}
			// 跳过重复
			for i < len(nums)-1 && nums[i] == nums[i+1] {
				i++
			}

		}
	}
	return res
}

// 四数之和
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	return nSumTarget(nums, 4, 0, int64(target))
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

func isPalindrome2(s string) bool {

	// 去除非数字和字母的部分
	var sb strings.Builder
	for _, r := range s {
		if unicode.IsDigit(r) || unicode.IsLetter(r) {
			sb.WriteByte(byte(unicode.ToLower(r)))
		}

	}
	s = sb.String()

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

// 867 转置矩阵
func transpose(matrix [][]int) [][]int {

	m := len(matrix)
	n := len(matrix[0])

	arr := make([][]int, n)
	for i := range arr {
		arr[i] = make([]int, m)
	}

	for i := range matrix {
		for j := range matrix[i] {
			arr[j][i] = matrix[i][j]
		}
	}
	return arr
}

// 151 反转字符串中的单词
func reverseWords(s string) string {

	split := strings.Split(s, " ")
	var res []string
	for i := range split {
		t := strings.TrimSpace(split[i])
		if t != "" {
			res = append(res, t)
		}
	}

	slices.Reverse(res)

	return strings.Join(res, " ")

}

// 48 旋转图像 n*n的矩阵
func rotate(matrix [][]int) {

	n := len(matrix)

	// 先按照左上到右下对角线把值翻转
	for i := 0; i < n; i++ {
		// 注意此处j=i,j的第一个值是在对角线上
		for j := i; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// 反转每一行
	for i := range matrix {
		row := matrix[i]

		left, right := 0, len(row)-1
		for left <= right {
			row[left], row[right] = row[right], row[left]
			left++
			right--
		}
		matrix[i] = row
	}
}

// 54 螺旋矩阵
func spiralOrder(matrix [][]int) []int {

	m := len(matrix)
	n := len(matrix[0])
	// 左上
	upperBound := 0
	// 右上
	rightBound := n - 1
	// 右下
	lowerBound := m - 1
	// 左下
	leftBound := 0

	var res []int
	for len(res) < m*n {

		if upperBound <= lowerBound {
			// 从右向左遍历
			for i := leftBound; i <= rightBound; i++ {
				res = append(res, matrix[upperBound][i])
			}
			upperBound++
		}

		// 向下遍历
		if leftBound <= rightBound {
			for i := upperBound; i <= lowerBound; i++ {
				res = append(res, matrix[i][rightBound])
			}
			rightBound--
		}

		// 向左遍历
		if upperBound <= lowerBound {
			for i := rightBound; i >= leftBound; i-- {
				res = append(res, matrix[lowerBound][i])
			}
			lowerBound--
		}

		// 向上遍历
		if leftBound <= rightBound {
			for i := lowerBound; i >= upperBound; i-- {
				res = append(res, matrix[i][leftBound])
			}
			leftBound++
		}

	}
	return res
}

// 59 螺旋矩阵II 根据n生成1~n*n的矩阵
func generateMatrix(n int) [][]int {

	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, n)
	}

	upperBound := 0
	rightBound := n - 1
	leftBound := 0
	lowerBound := n - 1

	for j := 1; j <= n*n; {

		// 从右向左遍历
		if upperBound <= lowerBound {
			for i := leftBound; i <= rightBound; i++ {
				res[upperBound][i] = j
				j++
			}
			upperBound++
		}

		// 向下遍历
		if leftBound <= rightBound {
			for i := upperBound; i <= lowerBound; i++ {
				res[i][rightBound] = j
				j++
			}
			rightBound--
		}

		// 向左遍历
		if upperBound <= lowerBound {
			for i := rightBound; i >= leftBound; i-- {
				res[lowerBound][i] = j
				j++
			}
			lowerBound--
		}

		// 向上遍历
		if leftBound <= rightBound {
			for i := lowerBound; i >= upperBound; i-- {
				res[i][leftBound] = j
				j++
			}
			leftBound++
		}

	}
	return res
}

// LCR 146. 螺旋遍历二维数组
func spiralArray(array [][]int) []int {
	var res []int
	if len(array) == 0 {
		return res
	}
	m := len(array)
	n := len(array[0])

	upperBound := 0
	rightBound := n - 1
	leftBound := 0
	lowerBound := m - 1

	for len(res) < m*n {
		// 向右遍历
		if upperBound <= lowerBound {
			for i := leftBound; i <= rightBound; i++ {
				res = append(res, array[upperBound][i])
			}
			upperBound++
		}

		// 向下遍历
		if leftBound <= rightBound {
			for i := upperBound; i <= lowerBound; i++ {
				res = append(res, array[i][rightBound])
			}
			rightBound--
		}

		// 向左遍历
		if upperBound <= lowerBound {
			for i := rightBound; i >= leftBound; i-- {
				res = append(res, array[lowerBound][i])
			}
			lowerBound--
		}

		// 向上遍历
		if leftBound <= rightBound {
			for i := lowerBound; i >= upperBound; i-- {
				res = append(res, array[i][leftBound])
			}
			leftBound++
		}

	}
	return res
}
