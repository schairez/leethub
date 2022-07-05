
// 435. Non-overlapping Intervals
// time: sort O(nlogn) + O(n) iter ≈ O(nlogn)
// space: O(logn) for sort recursion depth

func max(a, b int) int { if a >= b {return a}; return b}
func min(a, b int) int { if a <= b {return a}; return b}

func isOverlap(interval1, interval2 []int) bool {
    return max(interval1[0], interval2[0]) < min(interval1[1], interval2[1])
}

func eraseOverlapIntervals(intervals [][]int) int {
    n := len(intervals)
    res := 0
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][1] < intervals[j][1]
    })
    prev := intervals[0]
    for i := 1; i < n; i++ {
        if isOverlap(prev, intervals[i]) {
            res++
            continue
        }
        prev = intervals[i]
    }
    return res
}