/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
    if head == nil {
        return nil
    }
    
    it, res := head, head
    cnt := 0
    
    for cnt < k {
        cnt++
        if it.Next == nil {
            it.Next = head
            k = k%cnt
            cnt = 0
        }
        it = it.Next
    }
    
    for it.Next != head && it.Next != nil {
        it = it.Next
        res = res.Next
    }

    it.Next = head
    temp := res
    res = res.Next
    temp.Next = nil
    
    return res
}