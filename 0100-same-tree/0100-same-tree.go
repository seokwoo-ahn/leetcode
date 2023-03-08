/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
    if p == nil && q == nil {
        return true
    } else if p == nil || q == nil {
        return false
    } 
    check := true
    dfs(p, q, &check)
    return check
}

func dfs(p *TreeNode, q *TreeNode, check *bool) {
    if p.Val != q.Val {
        *check = false
        return
    }
    
    if p.Left != nil && q.Left != nil {
        dfs(p.Left, q.Left, check)        
    } else if p.Left == nil && q.Left == nil{
        
    } else if p.Left == nil || q.Left == nil {
        *check = false
        return         
    }
    
    if p.Right != nil && q.Right != nil {
         dfs(p.Right, q.Right, check)       
    } else if p.Right == nil && q.Right == nil{
        
    } else if p.Right == nil || p.Right == nil{
        *check = false
        return 
    } 
}
