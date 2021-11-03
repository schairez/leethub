import (
    "sort"
    "fmt"
)

//an array of intervals where intervals[i] = [starti, endi]


/*

[[1,4],[0,2],[3,5]]
sort -> [[0,2],[1,4],[3,5]]
res -> [[0,2]]

var (
    prevE, prevS
    currS, currE
)
res := [][]int{intervals[0]}
check := 0
for i:= 1; i < len(intervals); i++ {
    prevE, prevS = res[check][0], res[check][1] 
    currS, currE = intervals[i][0], intervals[i][1]    
    if currS > prevE {
        res = append(res, intervals[i])
    } else {       // else if currS <= prevE 
        res = append(res, min(prevS,))
    } 
}


*/

func merge(intervals [][]int) [][]int {
    //1 <= intervals.length <= 104
    if len(intervals) == 1 {
        return intervals
    }
    sort.Slice(intervals, func(i, j int) bool {
        if intervals[i][0] == intervals[j][0] {
            return intervals[i][1] < intervals[j][1]
        }
        return intervals[i][0] < intervals[j][0]
    })
    fmt.Println(intervals)
    res := [][]int{intervals[0]}
    fmt.Println(res)
    //prevS, prevE := res[0][0], res[0][1] 
    //check := 0
    for i:= 1; i < len(intervals); i++ {
        prevS, prevE := res[len(res)-1][0], res[len(res)-1][1]
        currS, currE := intervals[i][0], intervals[i][1]    
        if prevS <= currS && currE <= prevE { continue }
        if currS > prevE {
            res = append(res, intervals[i])
        } else {       // else if currS <= prevE 
            res[len(res)-1][1] = max(res[len(res)-1][1], currE)
        } 
    }
    return res
}

func min(a,b int) int {
    if a <= b  { return a }
    return b
}
func max(a,b int) int {
    if a >= b { return a }
    return b
}




/*
func merge(intervals [][]int) [][]int {
    //1 <= intervals.length <= 104
    if len(intervals) == 1 {
        return intervals
    }
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })
    res := [][]int{}
    fmt.Println(intervals)
    var currStart, currEnd int
    prevStart, prevEnd := intervals[0][0], intervals[0][1]
    res = append(res, intervals[0])
    for i:=1; i < len(intervals); i++ {
        currStart, currEnd = intervals[i][0], intervals[i][1]
        if prevStart -1 >= currStart || prevStart == 1 - currStart ||
        prevEnd >= 1 +  currStart  || prevEnd == currEnd || prevStart == currStart {
            res = res[1:]
            res = append(res, []int{min(prevStart,currStart), max(currStart, currEnd)})
        } else {
            res = append(res, intervals[i])
        } 
        prevStart, prevEnd = currStart, currEnd
    }
    return res
}
*/



/*
        if prevStart >= currStart {
            res = res[1:]
            res = append(res, []int{currStart, currEnd})
        } else if prevEnd >= currStart {
            res = res[1:]
            res = append(res, []int{prevStart, currEnd})
        } else {

*/