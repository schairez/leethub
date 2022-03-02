/*
           l               r = 8
 height = [1,8,6,2,5,4,8,3,7]
              w 
 49
*/

func maxArea(height []int) int {
    max := func(a, b int) int {
        if a >= b {return a}; return b
    }
    min := func(a, b int) int {
        if a <= b {return a}; return b
    }
    area := func(h, w int) int {
        return h * w
    }
    maxArea := 0
    n := len(height)
    start, end := 0, n-1 
    for start < end {
        w := end - start
        localArea := area(min(height[start], height[end]), w)  
        maxArea = max(maxArea, localArea)
        if height[start] <= height[end] {
            start++
        } else {
            end--
        }
    }
    return maxArea
}