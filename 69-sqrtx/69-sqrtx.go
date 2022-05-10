func mySqrt(x int) int {
    if x < 2 {
        return x
    }
    lo, hi := 1, x-1
    for lo + 1 < hi {
        mid := lo + (hi-lo) >> 1
        if mid*mid > x {
            hi = mid
        } else {
            lo = mid
        }
    }
    return lo

}


func mySqrtAlt(x int) int {
    if x <= 0 {
        return 0
    }
    lo, hi := 1, x-1
    for lo + 1 < hi {
        mid := lo + (hi-lo) >> 1
        if mid*mid <= x {
            lo = mid
        } else {
            hi = mid
        }
    }
    return lo
}

//time: O(logN)
//space:O(1)

func mySqrtV2(x int) int {
    target := x
    lo, hi := 1, target
    for lo < hi { //breaks when lo==hi
        m := lo + (hi-lo) / 2
        if m * m < target {
            lo = m+1
        } else {
            hi = m
        }
    }
    if lo * lo > target { 
        return lo - 1
    }
    
    return lo
}

