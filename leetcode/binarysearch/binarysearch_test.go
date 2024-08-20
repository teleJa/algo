package binarysearch

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSearch(t *testing.T) {

	nums := []int{-1, 0, 3, 5, 9, 12}
	t.Log(search(nums, 9))

}

func TestLeftBound(t *testing.T) {

	nums := []int{-1, 0, 3, 5, 7, 9, 12}
	t.Log(leftBound(nums, 8)) // 5

}

func TestRightBound(t *testing.T) {
	nums := []int{-1, 0, 3, 5, 7, 9, 12}
	t.Log(rightBound(nums, 8)) // 4
}

func TestSearchRange(t *testing.T) {

	tests := []struct {
		nums   []int
		target int
		want   []int
	}{
		{
			[]int{5, 7, 7, 8, 8, 10},
			8,
			[]int{3, 4},
		},
		{
			[]int{5, 7, 7, 8, 8, 10},
			7,
			[]int{1, 2},
		},
		{
			[]int{},
			0,
			[]int{-1, -1},
		},
	}

	for i, t0 := range tests {
		t.Run(fmt.Sprint("case", i), func(t *testing.T) {

			if !reflect.DeepEqual(searchRange(t0.nums, t0.target), t0.want) {
				t.Error("test fail")
			}

		})
	}

}

func Test_minEatingSpeed(t *testing.T) {
	type args struct {
		piles []int
		h     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{[]int{3, 6, 7, 11}, 8}, 4},
		{"case2", args{[]int{30, 11, 23, 4, 20}, 5}, 30},
		{"case3", args{[]int{30, 11, 23, 4, 20}, 6}, 23},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minEatingSpeed(tt.args.piles, tt.args.h); got != tt.want {
				t.Errorf("minEatingSpeed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_shipWithinDays(t *testing.T) {
	type args struct {
		weights []int
		days    int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5}, 15},
		{"case2", args{[]int{3, 2, 2, 4, 1, 4}, 3}, 6},
		{"case3", args{[]int{1, 2, 3, 1, 1}, 4}, 3},
	}

	for _, t0 := range tests {
		if w := shipWithinDays(t0.args.weights, t0.args.days); w != t0.want {
			t.Errorf("%s fail,w=%d,want=%d", t0.name, w, t0.want)
		}
	}

}

func Test_fnDays(t *testing.T) {
	type args struct {
		weights []int
		x       int
	}
	tests := []struct {
		name     string
		args     args
		wantDays int
	}{
		{"case1", args{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 14}, 5},
		{"case2", args{[]int{3, 2, 2, 4, 1, 4}, 6}, 3},
		{"case3", args{[]int{1, 2, 3, 1, 1}, 3}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotDays := fnDays(tt.args.weights, tt.args.x); gotDays != tt.wantDays {
				t.Errorf("fnDays() = %v, want %v", gotDays, tt.wantDays)
			}
		})
	}
}

func Test_searchMatrix(t *testing.T) {
	type args struct {
		matrix [][]int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"case1",
			args{[][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 3},
			true,
		},
		{
			"case2",
			args{[][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 13},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchMatrix(tt.args.matrix, tt.args.target); got != tt.want {
				t.Errorf("searchMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
