func topKFrequent(nums []int, k int) []int {
    sort.Ints(nums)
    
    q := &priorityQueue{}
	heap.Init(q)
    
    value := nums[0]
    cnt := 0
    for i := 0; i < len(nums); i++ {
        if value != nums[i] {
  		    heap.Push(q, element{value: value, count: cnt})
            value = nums[i]
            cnt = 1
        } else {
            cnt++
        }
        
        if i == len(nums) - 1 {
            heap.Push(q, element{value: value, count: cnt})
        }
    }

	ans := make([]int, k)
	for i := 0; i < k; i++ {
		v := heap.Pop(q)
		if s, ok := v.(element); ok {
			ans[i] = s.value
		}
	}
	return ans
}

type element struct {
	value int
	count int
}

type priorityQueue []element

func (q priorityQueue) Len() int           { return len(q) }
func (q priorityQueue) Less(i, j int) bool { return q[i].count > q[j].count }
func (q priorityQueue) Swap(i, j int)      { q[i], q[j] = q[j], q[i] }

func (q *priorityQueue) Push(x interface{}) {
	*q = append(*q, x.(element))
}

func (q *priorityQueue) Pop() interface{} {
	old := *q
	n := len(old)
	x := old[n-1]
	*q = old[0 : n-1]
	return x
}