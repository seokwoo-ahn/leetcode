func findHighAccessEmployees(access_times [][]string) []string {
    res := make([]string, 0)
    sort.Sort(CustomSort(access_times))
    
    if len(access_times) < 3 {
        return res
    }

    lastRes := ""
	for i := 0; i < len(access_times) - 2; i++ {
        if access_times[i][0] == access_times[i+2][0] && access_times[i][0] != lastRes {
            time1, _ := strconv.Atoi(access_times[i][1])
            time2, _ := strconv.Atoi(access_times[i+2][1])
            if time1 > time2 - 100 {
                res = append(res, access_times[i][0])
                lastRes = access_times[i][0]
            }
        }
	}
	return res
}

type CustomSort [][]string

func (cs CustomSort) Len() int {
	return len(cs)
}

func (cs CustomSort) Less(i, j int) bool {
	if cs[i][0] == cs[j][0] {
		return cs[i][1] < cs[j][1]
	}
	return cs[i][0] < cs[j][0]
}

func (cs CustomSort) Swap(i, j int) {
	cs[i], cs[j] = cs[j], cs[i]
}
