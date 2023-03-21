/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func swapPairs(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    if head.Next != nil{
        dummy := &ListNode {
            Val: 0,
            Next: &ListNode{},
        }
        dummy.Next = head.Next
        head.Next = swapPairs(head.Next.Next)
        dummy.Next.Next = head
        return dummy.Next
    } else {
        return head
    }
}