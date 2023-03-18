/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    q := []*TreeNode{root}
    depth := 0
    for len(q) > 0 {
        depth++
        num := len(q)
        for i:= 0; i < num; i++ {
            node := q[0]
            q = q[1:]
            if node.Left == nil && node.Right == nil {
                return depth
            }
            if node.Left != nil {
                q = append(q, node.Left)                
            }
            if node.Right != nil {
                q = append(q, node.Right)                
            }
        }
    }
    return depth
}