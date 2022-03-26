
    
// 1 <= len <= 10000
func canSeePersonsCount(heights []int) []int {
	n := len(heights)
	if n == 1 {
		return []int{0}
	}
	// mono decr
	var topIdx int
	var stack []int
	res := make([]int, n)
	/// 10, 6, when at 8
	for i := range heights {
		for {
			if len(stack) == 0 || heights[i] < heights[stack[len(stack)-1]] {
				break
			}
			topIdx, stack = stack[len(stack)-1], stack[:len(stack)-1]
			res[topIdx]++
			// pop elem
			// incr
		}
		if len(stack) != 0 {
			res[stack[len(stack)-1]]++
		}
		stack = append(stack, i)
	}
	return res
}
