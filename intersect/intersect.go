package intersect

func intersect(a []int, b []int) []int {
	m := make(map[int]int)
	for _, v := range a {
		m[v]++
	}
	arr := make([]int, 0)
	for _, v := range b {
		if m[v] > 0 {
			m[v]--
			arr = append(arr, v)
		}

	}
	return arr
}
