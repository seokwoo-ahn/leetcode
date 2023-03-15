/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
    res := 0;
    val := "";
    dfs(root, val, &res)
    return res
}

func dfs(node *TreeNode, val string, sum *int) {
    val = val + strconv.Itoa(node.Val)
    if node.Left == nil && node.Right == nil {
        value, _ := strconv.Atoi(val)
        *sum += value
        return 
    }
    
    if node.Left != nil {
        dfs(node.Left, val, sum)
    }
    
    if node.Right != nil {
        dfs(node.Right, val, sum)
    }
}
