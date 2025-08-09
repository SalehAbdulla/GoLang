package piscine

func BTreeLevelCount(root *TreeNode) int {
	if root == nil {return 0}
	getTotalLeft := BTreeLevelCount(root.Left)
	getTotalRight := BTreeLevelCount(root.Right)
	if getTotalLeft > getTotalRight {
		return getTotalLeft +1
	}
	return getTotalRight + 1
}
