package linkedlist

import (
	"log"
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {

	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 4}
	n1.Next = n2
	n2.Next = n3

	n4 := &ListNode{Val: 1}
	n5 := &ListNode{Val: 3}
	n6 := &ListNode{Val: 4}
	n4.Next = n5
	n5.Next = n6

	lists := mergeTwoLists(n1, n4)
	log.Println(lists)

}

func TestPartition(t *testing.T) {

	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 4}
	n1.Next = n2

	n3 := &ListNode{Val: 3}
	n2.Next = n3

	n4 := &ListNode{Val: 2}
	n3.Next = n4

	n5 := &ListNode{Val: 5}
	n4.Next = n5

	n6 := &ListNode{Val: 2}
	n5.Next = n6

	node := partition(n1, 3)
	log.Println(node)
}

func TestReverse(t *testing.T) {

	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 3}
	n1.Next = n2
	n2.Next = n3

	n := reverseN(n1, 2)
	log.Println(n)
}

func TestReverseN(t *testing.T) {

	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 3}
	n4 := &ListNode{Val: 4}
	n5 := &ListNode{Val: 5}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5

	n := reverseN(n1, 3)
	log.Println("node:", n)

}

func TestReverseBetween(t *testing.T) {

	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 3}
	n4 := &ListNode{Val: 4}
	n5 := &ListNode{Val: 5}
	n6 := &ListNode{Val: 6}
	n7 := &ListNode{Val: 7}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5
	n5.Next = n6
	n6.Next = n7

	node := reverseBetween(n1, 2, 6)
	log.Println("node:", node)

}

func TestReverseByNode(t *testing.T) {
	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 3}
	n4 := &ListNode{Val: 4}
	n5 := &ListNode{Val: 5}
	n6 := &ListNode{Val: 6}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5
	n5.Next = n6

	node := reverseByNode(n1, n4)
	log.Println("node:", node)

}

func TestRevereKGroup(t *testing.T) {
	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 3}
	n4 := &ListNode{Val: 4}
	n5 := &ListNode{Val: 5}
	n6 := &ListNode{Val: 6}
	n7 := &ListNode{Val: 7}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5
	n5.Next = n6
	n6.Next = n7

	node := reverseKGroup(n1, 3)
	log.Println(node)

}

func Test_isPalindrome(t *testing.T) {
	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 1}
	n3 := &ListNode{Val: 2}
	n4 := &ListNode{Val: 1}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	t.Log(isPalindrome(n1))

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
