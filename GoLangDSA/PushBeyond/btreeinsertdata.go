package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                 string
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	if root == nil {return &TreeNode{Data:data}}
	if data < root.Data {
		if root.Left == nil {
			root.Left = &TreeNode{Data:data, Parent: root}
		} else { // Go to left deeper until nil leaf reached
			BTreeInsertData(root.Left, data)
		}
	} else {
		if root.Right == nil {
			root.Right = &TreeNode{Data:data, Parent: root}
		} else { // Go deeper to the right until we reach nil leaf
			BTreeInsertData(root.Right, data)
		}
	}
	return root
}
