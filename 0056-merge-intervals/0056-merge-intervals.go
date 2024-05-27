func merge(intervals [][]int) [][]int {
    sort.Sort(CustomSort(intervals))
    
    res := make([][]int, 0)
    start, end := intervals[0][0], intervals[0][1]
    
    tempStart := intervals[0][0]
    
    for i := 1; i < len(intervals); i++ {
        if tempStart != intervals[i][0] {
            if intervals[i][0] > end {
                res = append(res, []int{start, end})
                start = intervals[i][0]
            }
            if end < intervals[i][1] {
                end = intervals[i][1]   
            }
            tempStart = intervals[i][0]
        }
    }
    
    res = append(res, []int{start, end})
    return res
}

type CustomSort [][]int

func (cs CustomSort) Len() int {
	return len(cs)
}

func (cs CustomSort) Less(i, j int) bool {
	if cs[i][0] == cs[j][0] {
		return cs[i][1] > cs[j][1]
	}
	return cs[i][0] < cs[j][0]
}

func (cs CustomSort) Swap(i, j int) {
	cs[i], cs[j] = cs[j], cs[i]
}
