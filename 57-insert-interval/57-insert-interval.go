
// ex: intervals = [[1,2], [3,5], [6,7]]; newInterval = [4,8]
// 
//
//
// ex: [[1,5]]; newInterval = [6,8]
//

func min(a, b int) int { if a <= b { return a}; return b}
func max(a, b int) int { if a >= b { return a}; return b}

func insert(intervals [][]int, newInterval []int) [][]int {
    n := len(intervals)
    if n == 0 {
        return [][]int{newInterval}
    }
    var stack [][]int
    var currStart, currEnd int
    var startTop, endTop int
    //check if we've merged newInterval
    isMerged := false
    idx := 0
    for idx != n {
        currStart, currEnd = intervals[idx][0], intervals[idx][1]
        if len(stack) != 0 {
            startTop, endTop = stack[len(stack)-1][0], stack[len(stack)-1][1]
            if !isMerged {
                if newInterval[0] <= endTop {
                    if newInterval[1] > endTop {
                        newStart := min(startTop, newInterval[0])
                        newEnd := max(endTop, newInterval[1])
                        stack[len(stack)-1][0], stack[len(stack)-1][1] = newStart, newEnd
                    }
                    isMerged = true
                } else {
                    if newInterval[0] < currStart {
                    stack = append(stack, newInterval)
                    isMerged = true
                    }
                }
            }
            startTop, endTop = stack[len(stack)-1][0], stack[len(stack)-1][1]
            if endTop < currStart {
                stack = append(stack, intervals[idx])
            } else {
                newStart := min(startTop, currStart)
                newEnd := max(endTop, currEnd)
                stack[len(stack)-1][0], stack[len(stack)-1][1] = newStart, newEnd
            }
        } else { //if stack is empty, not merged yet
            if currEnd < newInterval[0] { //newInterval starts at a later int val
                stack = append(stack, intervals[idx])
            } else {
                if newInterval[0] <= currEnd && newInterval[1] >= currStart {
                    newStart := min(currStart, newInterval[0])
                    newEnd := max(currEnd, newInterval[1])
                    stack = append(stack, []int{newStart, newEnd})
                } else {
                    stack = append(stack, newInterval)
                    stack = append(stack, intervals[idx])
                }
                isMerged = true
            }
        }
        idx++
    } 
    if !isMerged {
        if newInterval[0] > stack[len(stack)-1][1] {
            stack = append(stack, newInterval)
        } else {
            if newInterval[0] >= stack[len(stack)-1][0] {
                startTop, endTop := stack[len(stack)-1][0], stack[len(stack)-1][1]
                newStart := min(startTop, newInterval[0])
                newEnd := max(endTop, newInterval[1])
                stack[len(stack)-1][0], stack[len(stack)-1][1] = newStart, newEnd
            }
        }
    }
    return stack
}