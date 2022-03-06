//
//

func largestRectangleArea(heights []int) int {
    var maxArea int
    var stack []int 
    // in case we have mono increasing bars we'd end up
    // with our stack never popping
    heights = append(heights, 0)
    for i := range heights {
        for {
            if len(stack) == 0 || 
            heights[i] >= heights[stack[len(stack)-1]] {
                break
            }
            h := heights[stack[len(stack)-1]]
            stack = stack[:len(stack)-1]
            w := i
            if len(stack) > 0 {
                w = i - stack[len(stack)-1] - 1 
            }
            localArea := h * w
            if localArea > maxArea {
                maxArea = localArea
            }
        }
        stack = append(stack, i)
    }
    return maxArea
}