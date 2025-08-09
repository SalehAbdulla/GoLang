package piscine

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root == nil {return nil}
	if root.Data == elem {return root}

	var rootVal *TreeNode

	if elem < root.Data {
		rootVal = BTreeSearchItem(root.Left, elem)
	} else {
		rootVal = BTreeSearchItem(root.Right, elem)
	}
	
	return rootVal
}
