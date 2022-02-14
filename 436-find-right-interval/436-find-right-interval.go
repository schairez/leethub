
//436. Find Right Interval

//time: O(nlogn)
//space: O(n)

func findRightInterval(intervals [][]int) []int {
    n := len(intervals)
    if n == 1 {
        return []int{-1}
    }
    intervalData := make([][3]int, 0, n)
    for idx := range intervals {
        start, end := intervals[idx][0], intervals[idx][1]
        intervalData = append(intervalData, [3]int{start, end, idx})
    } 
    sort.Slice(intervalData, func(i, j int) bool {
        return intervalData[i][0] < intervalData[j][0]
    })
    res := make([]int, n)
    maxInt32 := 1 << 31 - 1
    for i :=0; i < n; i++ {
        idx := intervalData[i][2]
        endI := intervalData[i][1]
        val := -1
        min := maxInt32
        lo, hi := 0, n
        for lo < hi {
            mid := lo + (hi - lo) >> 1
            startJ := intervalData[mid][0]
            if startJ >= endI && startJ < min {
                hi = mid
                min = startJ
            } else {
                lo = mid+1
            } 
        }
        if min != maxInt32 {
            val = intervalData[lo][2]
        }
        res[idx] = val
    }
    return res
}

/*

[[3,4],[2,3],[1,2]]
[[3,4,0], [2,3,1], [1,2,2]]
 
*/





























//time: O(n^2)
//space: O(n)
//condition: each start := intervals[i][0] is unique
//startJ >= endI && startJ is minimized
/* 
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
*/