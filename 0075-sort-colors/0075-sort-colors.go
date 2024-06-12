func sortColors(nums []int)  {
    QuickSort(nums)
}

func QuickSort(arr []int) {
	if len(arr) < 2 {
		return
	}

	// 피벗 선택 및 분할 과정
	left, right := 0, len(arr)-1

	// 피벗을 중간값으로 설정합니다.
	pivotIndex := len(arr) / 2
	arr[pivotIndex], arr[right] = arr[right], arr[pivotIndex]

	// 피벗 기준으로 요소들을 재배치합니다.
	for i := range arr {
		if arr[i] < arr[right] {
			arr[i], arr[left] = arr[left], arr[i]
			left++
		}
	}

	// 피벗을 올바른 위치에 놓습니다.
	arr[left], arr[right] = arr[right], arr[left]

	// 피벗을 기준으로 왼쪽과 오른쪽 부분을 재귀적으로 정렬합니다.
	QuickSort(arr[:left])
	QuickSort(arr[left+1:])
}
