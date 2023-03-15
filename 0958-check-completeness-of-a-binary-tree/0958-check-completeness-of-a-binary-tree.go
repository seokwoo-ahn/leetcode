/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isCompleteTree(root *TreeNode) bool {
    queue := []*TreeNode{root}
    for len(queue) > 0 {
        q := queue[0]
        queue = queue[1:]
        
        if q == nil {
            break
        }
        
        queue = append(queue, q.Left, q.Right)
    }
    
    for _, q := range queue {
        if q != nil {
            return false
        }
    }
    
    return true
}