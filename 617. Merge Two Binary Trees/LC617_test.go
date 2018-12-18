/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

package lc617

import (
	"reflect"
	"testing"
	"fmt"
)

func Test_mergeTrees(t *testing.T) {
	type args struct {
		t1 *TreeNode
		t2 *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		// TODO: Add test cases.
		{
			name: "first",
			args:args{
				t1: arrayToTree([]int{1,3,2,5}),
				t2: arrayToTree([]int{2,1,3,0,4,0,7}),
			},
			want: arrayToTree([]int{3,4,5,5,4,0,7}),

		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTrees(tt.args.t1, tt.args.t2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTrees() = %v, want %v", got, tt.want)
			}
		})
	}
}

func arrayToTree(a []int) *TreeNode {
	var root *TreeNode
	for _, key := range a {
		root = insert(root, key)
	}

	inOrder(root, func(node *TreeNode) {
		fmt.Println(node.Val, node.Height)
	})
	return root
}

func leftRotate(root *TreeNode) *TreeNode {
	node := root.Right
	// fmt.Println(node.Val)
	root.Right = node.Left
	node.Left = root

	root.Height = max(height(root.Left), height(root.Right)) + 1
	node.Height = max(height(node.Right), height(node.Left)) + 1
	return node
}

func leftRigthRotate(root *TreeNode) *TreeNode {
	root.Left = leftRotate(root.Left)
	root = rightRotate(root)
	return root
}

func rightRotate(root *TreeNode) *TreeNode {
	node := root.Left
	root.Left = node.Right
	node.Right = root
	root.Height = max(height(root.Left), height(root.Right)) + 1
	node.Height = max(height(node.Left), height(node.Right)) + 1
	return node
}

func rightLeftRotate(root *TreeNode) *TreeNode {
	root.Right = rightRotate(root.Right)
	root = leftRotate(root)
	return root
}

func height(root *TreeNode) int {
	if root != nil {
		return root.Height
	}
	return -1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func insert(root *TreeNode, Val int) *TreeNode {
	if root == nil {
		root = &TreeNode{Val, 0, nil, nil}
		root.Height = max(height(root.Left), height(root.Right)) + 1
		return root
	}

	if Val < root.Val {
		root.Left = insert(root.Left, Val)
		if height(root.Left)-height(root.Right) == 2 {
			if Val < root.Left.Val {
				root = rightRotate(root) // 左左
			} else {
				root = leftRigthRotate(root) // 左右
			}
		}
	}

	if Val > root.Val {
		root.Right = insert(root.Right, Val)
		if height(root.Right)-height(root.Left) == 2 {
			if Val > root.Right.Val {
				root = leftRotate(root) // 右右
			} else {
				root = rightLeftRotate(root) // 右左
			}
		}
	}

	root.Height = max(height(root.Left), height(root.Right)) + 1
	return root
}

type action func(node *TreeNode)

func inOrder(root *TreeNode, action action) {
	if root == nil {
		return
	}

	inOrder(root.Left, action)
	action(root)
	inOrder(root.Right, action)
}
