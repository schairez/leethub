// 1984. Minimum Difference Between Highest and Lowest of K Scores
// time: O(nlogn) to sort
// space: O(logn) to sort + O(k) for queue space

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func min(a, b int) int { if a <= b { return a}; return b}

// calc min diff b/w highest and lowest scores
func calcMinDiff(nums []int) int {
    n := len(nums)
    lowest := nums[0]
    highest := nums[n-1]
    return abs(highest - lowest)
}

func minimumDifference(nums []int, k int) int {
    sort.Ints(nums)
    res := 1 << 31 - 1
    queue := make([]int, 0, k)
    for i := range nums {
        queue = append(queue, nums[i])
        if len(queue) == k {
            localMin := calcMinDiff(queue)
            res = min(res, localMin)
            queue = queue[1:]
        }
    }
    return res
}