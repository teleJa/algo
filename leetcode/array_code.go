package leetcode

import (
	"regexp"
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
