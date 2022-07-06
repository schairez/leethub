
func max(a, b int) int { if a >= b { return a}; return b}
func min(a, b int) int { if a <= b { return a}; return b}

func isOverlap(interval1 []int, interval2 []int) bool {
    return max(interval1[0], interval2[0]) <= min(interval1[1], interval2[1])
}

func merge(intervals [][]int) [][]int {
    n := len(intervals)
    res := make([][]int, 0, n)
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })
    prev := 0
    res = append(res, intervals[prev])
    for i := 1; i < n; i++ {
        if isOverlap(res[prev], intervals[i]) {
            res[prev][0] = min(res[prev][0], intervals[i][0]) 
            res[prev][1] = max(res[prev][1], intervals[i][1])
            continue
        }
        res = append(res, intervals[i])
        prev++
    }
    return res
}