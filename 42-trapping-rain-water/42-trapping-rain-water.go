func trap(height []int) int {
    var res int
    n := len(height)
    left, right := 0, n-1
    var maxLeft, maxRight int
    for left < right {
        if height[left] < height[right] {
            if height[left] >= maxLeft {
                maxLeft = height[left] 
            } else {
                res += maxLeft - height[left]
            }
            left++
        } else {
            if height[right] >= maxRight {
                maxRight = height[right]
            } else {
                res += maxRight - height[right]
            }
            right--
        }
    }
    return res
}

func min(a, b int) int { if a <= b { return a }; return b}

/*
func trap(height []int) int {
    res := 0
    var stack []int
    // time: O(n)
    // space: O(n)
    for idx := range height {
        for len(stack) > 0 && height[idx] > height[stack[len(stack)-1]] {
            bottom := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            if len(stack) > 0 {
                h := min(height[idx], height[stack[len(stack)-1]]) 
                w :=  idx - stack[len(stack)-1] - 1
                res += (h-height[bottom])*w
                fmt.Println(res)
            }
        }
        stack = append(stack, idx)
    }
    return res
}
*/