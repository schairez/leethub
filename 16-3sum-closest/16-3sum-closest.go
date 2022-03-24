
// 16. 3Sum Closest
// time: sort takes O(nlogn) + O(n^2) ≈ O(n^2)
// space: O(logn)

func min(a, b int) int {
    if a <= b {
        return a
    }
    return b
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func threeSumClosest(nums []int, target int) int {
    const maxInt32 = (1 << 31) - 1
    n := len(nums)
    if n < 3 {
        return 0
    }
    res := 0
    minDist := maxInt32
    sort.Ints(nums)
    for thirdIdx := n-1; thirdIdx > 1; thirdIdx-- {
        thirdVal := nums[thirdIdx]
        start, end := 0, thirdIdx-1
        for start < end {
            threeSum := nums[start] + nums[end] + thirdVal
            candDist := abs(target - threeSum)
            if candDist == 0 {
                return threeSum
            }
            if candDist < minDist {
               minDist = candDist
                res = threeSum
            } 
            if threeSum < target {
                start++
            } else {
                end--
            }
        }
    }
    return res
}
