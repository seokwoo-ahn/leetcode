/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    cur := head
    res := head
    
    for cur != nil {
        cur = cur.Next
        if n < 0 {
            res = res.Next
        }
        n--
    }
    if n == 0 {
        return head.Next
    }
    res.Next = res.Next.Next
    return head
}
