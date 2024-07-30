package leetcode

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
