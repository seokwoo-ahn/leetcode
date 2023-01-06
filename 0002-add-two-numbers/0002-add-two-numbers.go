/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    var sum int
    current := &ListNode{};
    result := current
    for l1 != nil || l2 != nil || sum != 0{
        if l1 != nil {
            sum += l1.Val
            l1 = l1.Next
        }
        
        if l2 != nil {
            sum += l2.Val
            l2 = l2.Next
        }
        
        value := sum % 10
        current.Next = &ListNode{value, nil}
        current = current.Next
        sum = sum/10
    }
    return result.Next
}