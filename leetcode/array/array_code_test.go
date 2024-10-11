package array

import (
	"log"
	"reflect"
	"regexp"
	"strings"
	"testing"
)

func TestMoveZeros(t *testing.T) {

	arr := []int{0, 1, 0, 3, 12}
	moveZeroes(arr)
	log.Println(arr)
}

func TestFilterLetterAndNumber(t *testing.T) {

	compile, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		t.Error("regex error:", err)
		return
	}

	allString := compile.ReplaceAllString("A man, a plan, a canal: Panama", "")
	t.Log(allString)
	s := "aA"

	t.Log(strings.EqualFold(string(s[0]), string(s[1])))

}

func TestLongestPalindrome(t *testing.T) {

	t.Log(longestPalindrome2("b"))
	t.Log(longestPalindrome2("bb"))
	t.Log(longestPalindrome2("cbbd"))

	t.Log(longestPalindrome("b"))
	t.Log(longestPalindrome("bb"))
	t.Log(longestPalindrome("cbbd"))

	t.Log(findPalindrome("cbbd", 1, 2))

}

func TestRemoveSliceFirst(t *testing.T) {

	arr := []int{1, 2, 3, 4, 5}
	arr = append(arr[1:])
	log.Println(arr)

	m := make(map[int][]int)
	ints := m[3]
	log.Println(len(ints))

	var a []int
	log.Println(len(a))
	log.Println(a == nil)

}

func Test_sortedSquares(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"case1", args{[]int{-4, -1, 0, 3, 10}}, []int{0, 1, 9, 16, 100}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortedSquares(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortedSquares() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortTransformedArray(t *testing.T) {
	type args struct {
		nums []int
		a    int
		b    int
		c    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"case1", args{[]int{-4, -2, 2, 4}, 1, 3, 5}, []int{3, 9, 15, 33}},
		{"case2", args{[]int{-4, -2, 2, 4}, -1, 3, 5}, []int{-23, -5, 1, 7}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortTransformedArray(tt.args.nums, tt.args.a, tt.args.b, tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortTransformedArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRotate(t *testing.T) {

	arr := []int{1, 2, 3, 4, 5}
	k := 3
	target := make([]int, len(arr))
	for i := range arr {
		target[(i+k)%len(arr)] = arr[i]
	}
	log.Println(target)
}

func TestJoin(t *testing.T) {

	split := []string{"hello", "world"}
	t.Log(strings.Join(split, "!"))

}

func Test_reverseWords(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"case1", args{"hello world"}, "world hello"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseWords(tt.args.s); got != tt.want {
				t.Errorf("reverseWords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_twoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"case2", args{[]int{3, 2, 4}, 6}, []int{1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveFromSlice(t *testing.T) {

	nums := []int{-1, -1, 0, 1, 2, 2}
	newNums := make([]int, len(nums[1:]))
	copy(newNums, nums[:2])
	copy(newNums[2:], nums[3:])
	log.Println(newNums) //  [-1 -1 1 2 2]
}

func Test_threeSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"case1", args{[]int{-1, 0, 1, 2, -1, -4}}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSum(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("threeSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
