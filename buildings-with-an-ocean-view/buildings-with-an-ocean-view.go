

/*
ocean view if all buildings to its right are smaller
*/



//1 <= heights[i] <= 10^9

//approach 1

//time: O(n)
//space: O(n)

func findBuildings(heights []int) []int {
    stack := make([]int, 0, len(heights))
    for idx, currHeight := range heights {
        for len(stack) > 0 && currHeight >= heights[stack[len(stack)-1]] {
            stack = stack[:len(stack)-1]
        } 
        stack = append(stack, idx)
    }
    return stack
}



//approach 2
//time: O(n)
//space: O(n)
func findBuildings2(heights []int) []int {
    maxHeight := 0  
    n := len(heights)
    res := make([]int, 0, n)
    for i := n-1; i >= 0; i-- {
        if heights[i] > maxHeight {
            res = append(res, i)
            maxHeight = heights[i]
        }
    }
    for start, end := 0, len(res)-1; start < end; start, end = start+1, end-1 {
        res[start], res[end] = res[end], res[start]
    }
    return res
}

