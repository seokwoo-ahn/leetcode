/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, targetSum int) bool {
    if root == nil {
        return false
    }
    return dfs(root, 0, targetSum)
}

func dfs(root *TreeNode, sum int, targetSum int) bool {
    sum = sum + root.Val
    if sum == targetSum && root.Left == nil && root.Right == nil {
        return true
    } else {
        if root.Left != nil {
            if dfs(root.Left, sum, targetSum) {
                return true
            }
        }
        
        if root.Right != nil {
            if dfs(root.Right, sum, targetSum) {
                return true
            }
        }
    }
    return false
}