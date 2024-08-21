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

// 410 分隔数组的最大值
func splitArray(nums []int, k int) int {
	return shipWithinDays(nums, k)
}

// 二维数组的二分搜索
// 二维数组的元素坐标可以抽象为一维数组坐标
// m行n列的二维数组,任意元素(i,j),展开为一维数组后坐标为 (i * n) + j
// 一维下标index也可以转换为二维坐标 (index/n,index % n)

// 74 搜索二维矩阵
// 每行中的整数从左到右按非严格递增顺序排列。
// 每行的第一个整数大于前一行的最后一个整数。
func searchMatrix(matrix [][]int, target int) bool {

	m := len(matrix)
	n := len(matrix[m-1])

	left, right := 0, (m-1)*n+n-1

	for left <= right {
		mid := left + (right-left)/2
		if fnMatrix(matrix, mid) == target {
			return true
		} else if fnMatrix(matrix, mid) < target {
			left = mid + 1
		} else if fnMatrix(matrix, mid) > target {
			right = mid - 1
		}

	}
	return false
}

func fnMatrix(matrix [][]int, mid int) int {
	m := len(matrix)
	n := len(matrix[m-1])
	return matrix[mid/n][mid%n]
}

// 240. 搜索二维矩阵 II
// 每行的元素从左到右升序排列
// 每列的元素从上到下升序排列
// 从右上角开始搜索,向左移动数据变小,向下移动数据变大
// 也可以从左下角开始搜索
func searchMatrix240(matrix [][]int, target int) bool {
	m := len(matrix) - 1
	n := len(matrix[0]) - 1

	i, j := 0, n
	for i <= m && j >= 0 {
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] < target {
			i++
		} else if matrix[i][j] > target {
			j--
		}
	}
	return false
}

// 566 重塑矩阵
// 展开为一维数组进行坐标映射即可
func matrixReshape(mat [][]int, r int, c int) [][]int {

	m := len(mat)
	n := len(mat[0])

	if r*c != m*n {
		return mat
	}
	res := make([][]int, r)
	for i := 0; i < r; i++ {
		res[i] = make([]int, c)
		// 填充元素
		for j := 0; j < len(res[i]); j++ {
			// 一维数组下标
			index := i*c + j
			res[i][j] = mat[index/n][index%n]
		}
	}
	return res
}

// 1,2     1,2,3
// 3,4  -> 4,5,6
// 5,6
