func largestRectangleArea(heights []int) int {
    // we append a 0 in the case that our heights input
    // is mono increasing, we'd never pop
    maxArea := 0
    heights = append(heights, 0)
    n := len(heights)
    stack := make([]int, 0, len(heights))
    for i := 0; i < n; i++ {
        for len(stack) != 0 && 
        heights[i] <= heights[stack[len(stack)-1]] {
            h := heights[stack[len(stack)-1]]
            stack = stack[:len(stack)-1]
            w := i
            if len(stack) != 0 {
                w = i - stack[len(stack)-1] - 1
            }
            if localArea := h * w; localArea > maxArea {
                maxArea = localArea
            }
        }
        stack = append(stack, i)
    }
    return maxArea
}