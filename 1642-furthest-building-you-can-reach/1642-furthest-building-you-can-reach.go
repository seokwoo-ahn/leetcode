func furthestBuilding(heights []int, bricks int, ladders int) int {
    h := &PriorityQueue{}
    heap.Init(h)
    
    if len(heights) == 1 {
        return 0
    }
    
    idx := 0
    sum := 0
    for idx < len(heights) - 1 {
        need := heights[idx + 1] - heights[idx]
        if need > 0 {
            sum += need
            heap.Push(h, need)
            if sum > bricks {
                if ladders > 0 {
                    ladders -= 1
                    val := heap.Pop(h).(int)
                    sum = sum - val
                } else {
                    break
                }
            }
        }
        idx++
    }
    
    return idx
}

type PriorityQueue []int

func (h PriorityQueue) Len() int { return len(h) }

func (h PriorityQueue) Less(i, j int) bool { return h[i] > h[j] }

func (h PriorityQueue) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *PriorityQueue) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *PriorityQueue) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
