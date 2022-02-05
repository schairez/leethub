
func abs(x int) int { if x < 0 { return -x }; return x }

//time: O(n)
//space: O(1)

//precondition: input nums is sorted in non-decr order
//so we can use two ptr technique

func sortedSquares(nums []int) []int {
    n := len(nums)
    res := make([]int, n)
    left, right := 0, n-1
    for idx := n-1; idx >=0; idx-- {
        if abs(nums[left]) >= abs(nums[right]) {
            res[idx] = nums[left] * nums[left]
            left++
        } else {
            res[idx] = nums[right] * nums[right]
            right--
        }
    }
    return res
}