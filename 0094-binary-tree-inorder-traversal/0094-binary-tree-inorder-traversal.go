/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    var res []int
    dfs(root, &res)
    return res
}

func dfs(root *TreeNode, res *[]int){
    if root != nil {
        dfs(root.Left, res)
        *res = append(*res, root.Val)
        dfs(root.Right, res)
    }
}