func validSubarrays(nums []int) int {
    var res int
    var stack []int
    for i := range nums {
        for {
            if len(stack) == 0 ||
            nums[i] >= nums[stack[len(stack)-1]] {
                break
            }
            stack = stack[:len(stack)-1]
        }
        stack = append(stack, i)
        res += len(stack)
    }
    return res
}