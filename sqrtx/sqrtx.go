//time: O(logN)
//space:O(1)


func mySqrt(x int) int {
    binSearch := func( target int) int {
    lo, hi := 1, target
    for lo < hi { //breaks when lo==hi
        m := lo + (hi-lo) / 2
        if m * m < target {
            lo = m+1
        } else {
            hi = m
        }
    }
    if lo * lo > target { return lo - 1}
    return lo
    }
    
    return binSearch(x)
    
}