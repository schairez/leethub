//time: O(nlogn + n) ~ O(nlogn)
//space: O(1)
import "sort"

func eraseOverlapIntervals(intervals [][]int) int {
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][1] < intervals[j][1]
    })
    n := len(intervals)
    cnt := 1
    end := intervals[0][1]
    for i:=1; i <len(intervals);i++ {
        if end <= intervals[i][0] {
            cnt++
            end = intervals[i][1]
        }
    }
    return n - cnt
    
}