/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func doubleIt(head *ListNode) *ListNode {
    carry := recursive(head)
    if carry != 0 {
        head = &ListNode{
            Val: carry,
            Next: head,
        }
    }
    
    return head
}

func recursive(node *ListNode) int {
    if node == nil {
        return 0
    }
    
    doubleVal := node.Val * 2 + recursive(node.Next)
    node.Val = doubleVal % 10
    
    return doubleVal / 10
}
