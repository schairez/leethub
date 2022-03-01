func min(a, b int) int { if a <= b {return a }; return b}

func trap(height []int) int {
    stack := make([]int, 0)
    var res int
    for idx := range height {
        for len(stack) > 0 && height[idx] > height[stack[len(stack)-1]] {
            n := len(stack)
            bottom := height[stack[n-1]]
            stack = stack[:n-1]
            if len(stack) > 0 {
                widthV := idx - stack[len(stack)-1] - 1
                heightV := min(height[stack[len(stack)-1]], height[idx])
                res += widthV * (heightV - bottom)
            }
        }
        stack = append(stack, idx)
    }
    return res
}