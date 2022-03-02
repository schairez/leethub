func min(a, b int) int {if a <= b { return a}; return b}

func trap(height []int) int {
    stack := make([]int, 0) 
    var res int
    for idx := range height {
        for len(stack) > 0 && height[idx] > height[stack[len(stack)-1]] {
            bottom := height[stack[len(stack)-1]]
            stack = stack[:len(stack)-1]
            if len(stack) > 0 {
                // 
                //w = idx - stack.top() - 1
                // 3 -1 -1  = 1
                w := idx - stack[len(stack)-1] - 1
                h := min(height[stack[len(stack)-1]], height[idx])
                res += w * (h - bottom)
            }
        }
        stack = append(stack, idx)
    }
    return res
}
// @ idx =3
// [1,]; bottom = ht[2] =0
// [0,1, 0] 2
// [1] 2; res += 
// height[stack.top()]

// bottom = 0