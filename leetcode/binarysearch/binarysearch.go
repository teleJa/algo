package binarysearch

// 寻找一个数
// 704 二分查找
func search(nums []int, target int) int {
	res := -1
	left, right := 0, len(nums)-1
	//搜索区间为空的时候退出循环
	for left <= right {
		// 直接(right+left) / 2 可能导致溢出
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		}
	}
	return res
}

// 寻找左侧边界
// 如果target不存在 返回大于target的最小索引
func leftBound(nums []int, target int) int {

	// 搜索区间为左闭右闭
	left, right := 0, len(nums)-1
	// 循环退出时 left = right + 1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			// 收缩右边界在[left,mid-1]中继续搜索
			right = mid - 1
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			// left增加,mid会右移即收缩左边界
			left = mid + 1
		}
	}

	if left < 0 || left >= len(nums) {
		return -1
	}

	// 未找到返回-1
	if nums[left] == target {
		return left
	}
	// 此处未找到返回-1,也可以直接返回left,此时的left即为大于target的最小索引
	return -1
}

// 寻找右侧边界
// 当target不存在时返回小于target的最大索引
func rightBound(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			// 增大左边界 在[mid+1,right]中继续搜索
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		}
	}

	if right < 0 || right >= len(nums) {
		return -1
	}

	if nums[right] == target {
		return right
	}
	// 未找到返回-1,也可以返回right
	return -1
}

// 34 寻找target的第一个和最后一个位置
func searchRange(nums []int, target int) []int {

	res := []int{-1, -1}

	// 搜索左边界
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			right = mid - 1
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		}
	}

	if left >= 0 && left < len(nums) && nums[left] == target {
		res[0] = left
	}
	// 搜索右边界
	left, right = 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			left = mid + 1
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		}
	}

	if right >= 0 && right < len(nums) && nums[right] == target {
		res[1] = right
	}
	return res
}

// 875 吃香蕉
func minEatingSpeed(piles []int, h int) int {

	left, right := 1, 1000000000
	for left <= right {
		mid := left + (right-left)/2
		if f(piles, mid) == int64(h) {
			right = mid - 1
		} else if f(piles, mid) > int64(h) {
			left = mid + 1
		} else if f(piles, mid) < int64(h) {
			right = mid - 1
		}
	}

	return left
}

func f(piles []int, x int) (h int64) {

	for _, i := range piles {
		h += int64(i / x)
		if i%x > 0 {
			h++
		}
	}
	return
}

// 1011
func shipWithinDays(weights []int, days int) int {
	var left, right int
	for i := range weights {
		if weights[i] > left {
			left = weights[i]
		}
		right += weights[i]
	}

	for left <= right {
		mid := left + (right-left)/2
		if fnDays(weights, mid) == days {
			right = mid - 1
		} else if fnDays(weights, mid) > days {
			left = mid + 1
		} else if fnDays(weights, mid) < days {
			right = mid - 1
		}
	}
	return left
}

func fnDays(weights []int, x int) (days int) {
	for i := 0; i < len(weights); {
		cap := x
		for i < len(weights) {
			if cap < weights[i] {
				break
			} else {
				cap -= weights[i]
			}
			i++
		}
		days++

	}
	return
}
