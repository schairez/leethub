// 239. Sliding Window Maximum
// mono decr queue approach
// time: O(n) 
// space: O(k) for the deque + O(n-k+1) for res

func maxSlidingWindow(nums []int, k int) []int {
    n := len(nums)
    numWin := n - k + 1
    res := make([]int, 0, numWin)
    deque := make([]int, 0, k)
    
    for i := 0; i < n; i++ {
        if len(deque) != 0 && deque[0] == i-k {
            deque = deque[1:]
        }
        for {
            if len(deque) == 0 || nums[deque[len(deque)-1]] >= nums[i] {
                break
            }
            deque = deque[:len(deque)-1]
        }
        deque = append(deque, i)
        if i >= k - 1 {
            res = append(res, nums[deque[0]])
        }
    }
    return res
}