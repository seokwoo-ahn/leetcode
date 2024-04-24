func findKthLargest(nums []int, k int) int {
    var heap MaxHeap
    heap = make([]int, 0)
    
    for i := 0; i < len(nums); i++ {
        (&heap).push(nums[i])
    }
    
    res := 0
    for i := 0; i < k; i++ {
        res = (&heap).pop()
    }
    return res
}

type MaxHeap []int

func (h *MaxHeap) push(v int) {
    *h = append(*h, v)
    index := len(*h) - 1

    for index > 0 {
        parentIndex := (index - 1) / 2
        if (*h)[parentIndex] >= (*h)[index] {
            break
        }
        (*h)[parentIndex], (*h)[index] = (*h)[index], (*h)[parentIndex]
        index = parentIndex
    }
}

func (h *MaxHeap) pop() (int) {
    if len(*h) == 0 {
        return 0
    }

    res := (*h)[0]
    lastIndex := len(*h) - 1
    (*h)[0] = (*h)[lastIndex]
    (*h) = (*h)[:lastIndex]

    index := 0
    for {
        leftChildIndex := 2*index + 1
        rightChildIndex := 2*index + 2
        largest := index

        if leftChildIndex < len(*h) && (*h)[leftChildIndex] > (*h)[largest] {
            largest = leftChildIndex
        }
        if rightChildIndex < len(*h) && (*h)[rightChildIndex] > (*h)[largest] {
            largest = rightChildIndex
        }
        if largest == index {
            break
        }

        (*h)[index], (*h)[largest] = (*h)[largest], (*h)[index]
        index = largest
    }

    return res
}