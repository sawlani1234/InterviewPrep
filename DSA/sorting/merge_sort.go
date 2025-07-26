package sorting

type MergeSort struct {
}

func (m *MergeSort) Merge(a []int, b []int) []int {
	left, right := 0, 0

	ans := make([]int, 0)
	for left < len(a) && right < len(b) {
		if a[left] < b[right] {
			ans = append(ans, a[left])
			left++
		} else {
			ans = append(ans, b[right])
			right++
		}
	}

	for ; left < len(a); left++ {
		ans = append(ans, a[left])
	}

	for ; right < len(b); right++ {
		ans = append(ans, b[right])
	}
	return ans
}

func (m *MergeSort) RecursiveMerge(a []int) []int {
	if len(a) == 1 {
		return a
	}

	mid := len(a) / 2
	left := m.RecursiveMerge(a[:mid])
	right := m.RecursiveMerge(a[mid:])
	return m.Merge(left, right)
}

func (m *MergeSort) Sort(a []int) []int {
	return m.RecursiveMerge(a)
}
