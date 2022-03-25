
//import "math/rand"

func init() {
    rand.Seed(time.Now().UnixNano())
}


// iterative quick select approach
// random partition using boyer-yates algorithm shuffle

func findKthLargest(nums []int, k int) int {
    n := len(nums)
    rand.Shuffle(n, func(i, j int) {
        nums[i], nums[j] = nums[j], nums[i]
    })
    
    kIdx := n - k
    for l, r := 0, n-1; l < r; {
        pivotElem := nums[l]
        start, end := l, r+1
        for {
            for start++; start < r && nums[start] <= pivotElem; start++ {}
            for end--; end > l && pivotElem <= nums[end]; end-- {}
            if start >= end {
                break
            }
            nums[start], nums[end] = nums[end], nums[start]
        }
        nums[l], nums[end] = nums[end], pivotElem
        if end == kIdx {
            break
        } else if end < kIdx {
            l = end + 1
        } else {
            r = end - 1
        }
    }
    return nums[kIdx]
}