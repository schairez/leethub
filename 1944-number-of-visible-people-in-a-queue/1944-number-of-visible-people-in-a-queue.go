func canSeePersonsCount(heights []int) []int {
    n := len(heights)
    var stack []int
    res := make([]int, n)
    for i := range heights {
        for {
            if len(stack) == 0 ||
            heights[stack[len(stack)-1]] > heights[i] {
                break
            } 
            idx := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            res[idx]++
        }
        if len(stack) > 0 {
            res[stack[len(stack)-1]]++
        }
        stack = append(stack, i)
    }
    return res
}