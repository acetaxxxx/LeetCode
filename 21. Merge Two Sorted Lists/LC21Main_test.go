/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

package LeetCode

import (
	"github.com/stretchr/testify/assert"

	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		l1 []int
		l2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "test",
			args: args{
				l1: []int{1, 2, 4},
				l2: []int{1, 3, 4},
			},
			want: []int{1, 1, 2, 3, 4, 4},
		},
		{
			name: "test2",
			args: args{
				l1: []int{1, 2, 4, 4, 4, 4},
				l2: []int{1, 3, 4, 6, 6, 64, 100},
			},
			want: []int{1, 1, 2, 3, 4, 4, 4, 4, 4, 6, 6, 64, 100},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l1 := makeListNode(tt.args.l1)
			l2 := makeListNode(tt.args.l2)
			want := makeListNode(tt.want)
			got := mergeTwoLists(l1, l2)
			assert.Equal(t, ListToArray(want), ListToArray(got))

		})
	}
}
func makeListNode(a []int) *ListNode {

	var result *ListNode

	if len(a) == 0 {
		return result
	}
	result = new(ListNode)
	result.Val = a[0]
	iter := result
	for i := 1; i < len(a); i++ {
		iter.Next = new(ListNode)
		iter = iter.Next
		iter.Val = a[i]
	}
	return result
}
func ListToArray(l *ListNode) []int {
	var r []int
	for l.Next != nil {
		r = append(r, l.Val)
		l = l.Next
	}
	return r

}
