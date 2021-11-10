import "sort"
//time: O(nlogn)
//space: O(logN)

//an array of intervals where intervals[i] = [starti, endi]

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
    res := [][]int{intervals[0]}
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

func max(a,b int) int {
    if a >= b { return a }
    return b
}


