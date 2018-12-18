/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
package lc617

type TreeNode struct {
	Val    int
	Height int
	Left   *TreeNode
	Right  *TreeNode
	
}

type oper struct {
	t1     *TreeNode
	t2     *TreeNode
	result *TreeNode
}

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	var queue []oper

	if t1 == nil && t2 == nil {
		return nil
	}
	result := new(TreeNode)
	queue = append(queue, oper{
		t1:     t1,
		t2:     t2,
		result: result,
	})
	for len(queue) > 0 {
		tmp := queue[0]
		queue = mergeNode(tmp.t1, tmp.t2, tmp.result, queue)
		queue = queue[1:]
	}

	return result
}

func mergeNode(t1 *TreeNode, t2 *TreeNode, result *TreeNode, queue []oper) []oper{
	if t1 == nil && t2 == nil { // 左右均為nil
		return queue
	} else if t1 == nil { //僅左邊為nil
		result.Val = t2.Val
		if t2.Left != nil {
			result.Left = new(TreeNode)
			queue = append(queue, oper{
				t1:     t1,
				t2:     t2.Left,
				result: result.Left,
			})
		}
		if t2.Right != nil {
			result.Right = new(TreeNode)
			queue = append(queue, oper{
				t1:     t1,
				t2:     t2.Right,
				result: result.Right,
			})
		}
	} else if t2 == nil { // 僅右邊為nil

		result.Val = t1.Val

		if t1.Left != nil {
			result.Left = new(TreeNode)
			queue = append(queue, oper{
				t1:     t1.Left,
				t2:     t2,
				result: result.Left,
			})
		}
		if t1.Right != nil {
			result.Right = new(TreeNode)
			queue = append(queue, oper{
				t1:     t1.Right,
				t2:     t2,
				result: result.Right,
			})
		}
	} else { //兩邊均不為nil
		result.Val = t1.Val + t2.Val

		if t1.Left != nil || t2.Left != nil {
			result.Left = new(TreeNode)

			queue = append(queue, oper{
				t1:     t1.Left,
				t2:     t2.Left,
				result: result.Left,
			})
		}
		if t1.Right != nil || t2.Right != nil {
			result.Right = new(TreeNode)

			queue = append(queue, oper{
				t1:     t1.Right,
				t2:     t2.Right,
				result: result.Right,
			})
		}
	}
	return queue
}
