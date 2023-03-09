/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDiffInBST(root *TreeNode) int {
    values := make([]int, 0)
    dfs(root, &values)
    sort.Ints(values)
    result := math.MaxInt
    
    for i:= 0; i < len(values) -1; i++ {
        tempValue := values[i+1]-values[i]
        if tempValue == 0 {
            return 0
        }
        
        if result > tempValue {
            result = tempValue
        }
    }
    return result
}

func dfs(node *TreeNode, values *[]int) {
    *values = append(*values, node.Val)
    
    if node.Left != nil {
        dfs(node.Left, values)
    }
    
    if node.Right != nil {
        dfs(node.Right, values)
    }
}

