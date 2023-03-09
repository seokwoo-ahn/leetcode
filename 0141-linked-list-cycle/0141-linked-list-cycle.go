/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
    check := make(map[*ListNode]bool)
    
    for head != nil {
        if check[head] == true {
            return true
        }
        check[head] = true
        head = head.Next
    }
    return false
}