/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func longestZigZag(root *TreeNode) int {
    if root == nil {
        return 0
    }
    res := 0
    if root.Left != nil {
        dfs(root.Left, true, 1, &res)        
    }
    
    if root.Right != nil {
        dfs(root.Right, false, 1, &res)   
    }
    return res
}

func dfs(node *TreeNode, left bool, tmp int, res *int) {
    if !left {
        if node.Left != nil {
            dfs(node.Left, true, tmp+1, res)
        } else {
            if *res < tmp {
                *res = tmp
            }
        }
        
        if node.Right != nil {
            dfs(node.Right, false, 1, res)
        }
    } else {
        if node.Right != nil {
            dfs(node.Right, false, tmp+1, res)           
        } else {
            if *res < tmp {
                *res = tmp
            }
        }
        
        if node.Left != nil {
            dfs(node.Left, true, 1, res)
        }
    }
}