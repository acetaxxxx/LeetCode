package lc687

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Mark
type Mark struct {
	Val    int
	Length int
}

// newTreeNode
type NodeResult struct {
	MaxChild *Mark
	SelfNode *Mark
}

func longestUnivaluePath(root *TreeNode) int {

	if root == nil {
		return 0
	}
	
}

func getLongestPath(node *TreeNode)*NodeResult{
	
	if  node == nil {
		return nil
	}

	leftResult := getLongestPath(node.Left)
	rightResult := getLongestPath(node.Right)

	nodeResult := new(NodeResult)
	
	
}
