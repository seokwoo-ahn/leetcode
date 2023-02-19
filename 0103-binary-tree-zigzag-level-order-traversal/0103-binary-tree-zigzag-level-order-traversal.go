/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Node struct {
    TreeNode
    Depth int    
}

func zigzagLevelOrder(root *TreeNode) [][]int {
    var queue []*Node
    var result [][]int
    
    if root == nil {
        return nil
    }
    
    rootNode := &Node{
        TreeNode: *root,
        Depth: 0,
    }
    
    queue = append(queue, rootNode)
    
    var currentDepth int
    var currentResult []int
    isLeft2Right := true
    
    for len(queue) != 0 {
        node := queue[0]
        queue = queue[1:]
        
        if currentDepth != node.Depth {
            currentDepth = node.Depth
            result = append(result, currentResult)
            currentResult = []int{}
            isLeft2Right = !isLeft2Right
        }
        
        if isLeft2Right {
            currentResult = append(currentResult, node.Val)
        } else {
            currentResult = append([]int{node.Val}, currentResult...)
        } 

        
        if node.Left != nil {
            newNode := &Node{
                TreeNode: *node.Left,
                Depth: node.Depth+1,
            }
            
            queue = append(queue, newNode)
        }
        
        if node.Right != nil {
            newNode := &Node{
                TreeNode: *node.Right,
                Depth: node.Depth+1,
            }
            
            queue = append(queue, newNode)
        }
        
    }
    result = append(result, currentResult)
    return result
}