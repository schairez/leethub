/*
ocean view if all buildings to its right are smaller
*/

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