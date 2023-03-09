/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
    check := make(map[*ListNode]bool)
    
    for head != nil {
        if check[head] == true {
            return head
        }
        check[head] = true
        head = head.Next
    }
    return nil
}