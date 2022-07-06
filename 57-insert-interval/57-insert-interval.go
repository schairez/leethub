func max(a, b int) int { if a >= b { return a}; return b}
func min(a, b int) int { if a <= b { return a}; return b}


func isOverlap(interval1 []int, interval2 []int) bool {
    return max(interval1[0], interval2[0]) <= min(interval1[1], interval2[1])
}

func insert(intervals [][]int, newInterval []int) [][]int {
    // condition: input presorted by start time
    n := len(intervals)
    if n == 0 {
        return [][]int{newInterval}
    }
    res := make([][]int, 0, n+1)
    i := 0
    for i < n {
        if isOverlap(intervals[i], newInterval) {
            newInterval[0] = min(intervals[i][0], newInterval[0])
            newInterval[1] = max(intervals[i][1], newInterval[1])
        } else if intervals[i][1] < newInterval[0] {
            res = append(res, intervals[i])
        } else {
            break
        }
        i++
    }
    res = append(res, newInterval)
    for i < n {
        res = append(res, intervals[i])
        i++
    }
    return res
}