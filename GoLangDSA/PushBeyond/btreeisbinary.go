package piscine

func BTreeIsBinary(root *TreeNode) bool {
	if root == nil {return true}
	if root.Data > root.Left.Data && root.Data < root.Right.Data {
		return true
	} else {
		return false
	}
}
