import "sort"

func canAttendMeetings(intervals [][]int) bool {
    if len(intervals) == 0 { return true }
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][1] < intervals[j][1] 
    })
    
    var currStart, currEnd int
    prevStart, prevEnd := intervals[0][0], intervals[0][1]
    for i:=1; i < len(intervals); i++ {
        currStart, currEnd = intervals[i][0], intervals[i][1]
        if prevStart > currStart || prevEnd > currStart || prevStart > currEnd { 
            return false 
        } 
        prevStart, prevEnd = currStart, currEnd
    }
    return true
}

/*

[[16,22],[28,45],[3,9],[46,50],[13,14]]




*/

