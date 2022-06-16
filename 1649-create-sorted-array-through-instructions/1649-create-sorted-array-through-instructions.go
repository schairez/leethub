
func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func createSortedArray(instructions []int) int {
    const mod = 1_000_000_007
	nums := make([]int, len(instructions))
	cost := 0

	for length, num := range instructions {
		l, r := sort.Search(length, func(i int) bool { return nums[i] >= num }), sort.Search(length, func(i int) bool { return nums[i] > num })

		cost += min(l, length-r)

		copy(nums[r+1:length+1], nums[r:length+1])
		nums[r] = num
	}
	return cost % mod
}