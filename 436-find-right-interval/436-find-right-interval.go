//time: O(n^2)
//space: O(n)

//condition: each start := intervals[i][0] is unique
//startJ >= endI && startJ is minimized


func findRightInterval(intervals [][]int) []int {
    n := len(intervals)
    if n == 1 {
        return []int{-1}
    }
    //map interval to idx of input
    //arrays are comparable in go
    intervalIdxMap := make(map[[2]int]int, n)
    for idx, interval := range intervals {
        start, end := interval[0], interval[1]
        intervalIdxMap[[2]int{start,end}] = idx
    }
    sort.Slice(intervals, func(i, j int) bool {
        //sort by start
        return intervals[i][0] < intervals[j][0]
    })
    res := make([]int, n)
    for i := 0; i < len(intervals); i++ {
        startI, endI := intervals[i][0], intervals[i][1]
        min := 1 << 31 - 1
        val := -1
        for j := i; j < n; j++ {
            if intervals[j][0] >= intervals[i][1] &&
            intervals[j][0] < min {
                min = intervals[j][0]
                startJ, endJ := intervals[j][0], intervals[j][1]
                val = intervalIdxMap[[2]int{startJ,endJ}]
            }
        }
        idx := intervalIdxMap[[2]int{startI,endI}]
        res[idx] = val
    }
    return res
}