/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }
    
    dfs(root)

    return root
}

func dfs(node *TreeNode) {
    if node.Left == nil && node.Right == nil {
        return
    }
    invert(node)
    
    if node.Left != nil {
        dfs(node.Left)
    }
    
    if node.Right != nil{
        dfs(node.Right)
    }
}

func invert(node *TreeNode) {
    temp := node.Left
    node.Left = node.Right
    node.Right = temp
}