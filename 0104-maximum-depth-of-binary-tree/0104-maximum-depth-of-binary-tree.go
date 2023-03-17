/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0;
    }
    res := 0
    dfs(root, 1, &res)
    return res
}

func dfs(node *TreeNode, curDepth int, maxDepth *int) {
    if curDepth > *maxDepth {
        *maxDepth = curDepth
    }
    
    if node.Left != nil {
        dfs(node.Left, curDepth + 1, maxDepth)
    }
    
    if node.Right != nil {
        dfs(node.Right, curDepth + 1, maxDepth)
    }
}