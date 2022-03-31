// 228. Summary Ranges
// time: O(n)
// space: O(n)
func summaryRanges(nums []int) []string {
    n := len(nums)
    if n == 0 {
        return []string{}
    }
    res := make([]string, 0, n)
    start := nums[0]
    var end int
    //"a->b" if a != b
    //"a" if a == b
    strRangeAppend := func(a, b int) {
        if a != b {
            res = append(res, fmt.Sprintf("%d->%d",start, end))
        } else {
            res = append(res, fmt.Sprintf("%d", start))
        }
    }
    for i := 0; i < n-1; i++ {
        if nums[i] + 1 != nums[i+1] {
            end = nums[i]
            strRangeAppend(start, end)
            start = nums[i+1]
        }
    }
    end = nums[n-1]
    strRangeAppend(start, end)
    
    return res
}