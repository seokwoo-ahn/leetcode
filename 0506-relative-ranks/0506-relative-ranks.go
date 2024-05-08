func findRelativeRanks(score []int) []string {
    h := &Heap{}
    
    heap.Init(h)
    
    for i := 0; i < len(score); i++ {
        heap.Push(h, Element{Val: score[i], Idx: i})
    }
    
    res := make([]string, len(score))
    for i := 0; i < len(score); i++ {
        if i == 0 {
            res[heap.Pop(h).(int)] = "Gold Medal"
        } else if i == 1 {
            res[heap.Pop(h).(int)] = "Silver Medal"
        } else if i == 2 {
            res[heap.Pop(h).(int)] = "Bronze Medal"
        } else {
            res[heap.Pop(h).(int)] = strconv.Itoa(i + 1)
        }
    }
    return res
}

type Heap []Element

type Element struct {
    Val int
    Idx int
}

func (h Heap) Len() int {
	return len(h)
}

func (h Heap) Less(i, j int) bool {
	return h[i].Val > h[j].Val
}

func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *Heap) Push(element interface{}) {
	*h = append(*h, element.(Element))
}

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	element := old[n-1]
	*h = old[0 : n-1]
	return element.Idx
}
