/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type Solution struct {
    ListNode
    Head *ListNode
}


func Constructor(head *ListNode) Solution {
    return Solution{*head, head}
}


func (this *Solution) GetRandom() int {
    randNum := rand.Intn(10000)
    cur := this.Head
    count := 1
    for randNum > 0 {
        if cur.Next == nil {
            cur.Next = this.Head
            randNum = randNum%count
        }
        randNum--
        count++
        cur = cur.Next
    }
    return cur.Val
}



/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */
