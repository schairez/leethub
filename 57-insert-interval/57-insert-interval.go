// 57. Insert Interval
// time: O(n)
// space: O(n)

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
    idx := 0
    for idx < n {
        if isOverlap(intervals[idx], newInterval) {
            newInterval[0] = min(intervals[idx][0], newInterval[0])
            newInterval[1] = max(intervals[idx][1], newInterval[1])
        } else if intervals[idx][1] < newInterval[0] {
            res = append(res, intervals[idx])
        } else {
            break
        }
        idx++
    }
    res = append(res, newInterval)
    for idx < n {
        res = append(res, intervals[idx])
        idx++
    }
    return res
}