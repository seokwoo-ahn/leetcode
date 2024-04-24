func findKthLargest(nums []int, k int) int {
    heap := make([]int, 0)
    
    for i := 0; i < len(nums); i++ {
        push(&heap, nums[i])
    }
    
    res := 0
    for i := 0; i < k; i++ {
        res = pop(&heap)
    }
    return res
}

func push(heap *[]int, v int) {
    idx := len(*heap)
    *heap = append(*heap, v)
    convertMaxHeapAfterPush(heap, idx)
}

func convertMaxHeapAfterPush(heap *[]int, idx int) { // 포인터로 변경
    if idx == 0 {
        convertMaxHeapAfterPop(heap, idx)
        return
    }
    
    parent := (idx-1)/2
    child := idx
    
    if (*heap)[parent] < (*heap)[child] {
        swap(&(*heap)[parent], &(*heap)[child])
        convertMaxHeapAfterPush(heap, parent)
    }
}

func pop(heap *[]int) int {
    v := (*heap)[0]
    swap(&(*heap)[0], &(*heap)[len(*heap) - 1])
    *heap = (*heap)[:(len(*heap) - 1)]
    convertMaxHeapAfterPop(heap, 0)
    return v
}

func convertMaxHeapAfterPop(heap *[]int, idx int) { // 포인터로 변경
    parent := idx
    leftChild := idx*2 + 1 // 수정
    rightChild := idx*2 + 2 // 수정
    
    if leftChild > len(*heap) - 1 {
        return
    } else {
        maxChild := leftChild
        if rightChild <= len(*heap) - 1  && (*heap)[rightChild] > (*heap)[leftChild] {
            maxChild = rightChild
        }
        
        if (*heap)[parent] < (*heap)[maxChild] {
            swap(&(*heap)[parent], &(*heap)[maxChild])
            convertMaxHeapAfterPop(heap, maxChild)
        }
    }
}

func swap(a *int, b *int) {
    temp := *a
    *a = *b
    *b = temp
}
