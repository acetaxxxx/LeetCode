package lc687

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type NewTreeNode struct{
	Left *NewTreeNode
	Right *NewTreeNode
	Val int
	Length int
	Parent *NewTreeNode
}

type Marked struct {
	Val    int
	Length int
}

func longestUnivaluePath(root *TreeNode) int {

	if root == nil {
		return 0
	}
	
	newTree := createNewTreeNode(root)


}

func createNewTreeNode(root *TreeNode ) *NewTreeNode {

	if root == nil {return nil}

	node := &NewTreeNode{
		Val: root.Val,
		Length: 0,
	}

}


func getMarkNode(node *TreeNode, parrentMark *Marked) *Marked {
	if node == nil {
		return nil
	}
	nodeMark := &Marked{
		Val:    node.Val,
		Length: 0,
	}

	if parrentMark != nil && parrentMark.Val == node.Val {
		nodeMark.Length = parrentMark.Length + 1
	}

	leftMax := getMarkNode(node.Left, nodeMark)
	rightMax := getMarkNode(node.Right, nodeMark)

	return getLargestMark(leftMax, rightMax, nodeMark)

}

func getLargestMark(left *Marked, right *Marked, node *Marked) *Marked {
	if left == nil && right == nil {
		return node
	}

	if left == nil {
		if right.Length > node.Length {
			return right
		}
		return node
	}
	if right == nil {
		if left.Length > node.Length {
			return left
		}
		return node
	}

	// node is max
	if node.Length >= right.Length && node.Length >= left.Length {
		return node
	}

	if left.Length > right.Length {
		return left
	}
	return right
}
