func binSearch(arr []int, target int) int {
    n := len(arr)
    lo, hi := 0, n-1
    for lo + 1 < hi {
        mid := lo + (hi-lo) >> 1
        if arr[mid] >= target {
            hi = mid
        } else {
            lo = mid
        }
    }
    if arr[lo] >= target {
        return arr[lo]+1
    } else if arr[hi] >= target {
        return arr[hi]+1
    }
    return arr[n-1]+1 
}

/*
src ="xyz"
t = "xzyxz"

res, curr = 0, 0
curr = 2







*/

func shortestWay(source string, target string) int {
    var charIdxMap [26][]int
    for idx := range source {
        charIdxMap[source[idx]-'a'] = append(charIdxMap[source[idx]-'a'], idx)
    }
    res, currIdx := 1, -1
    for idx := range target {
        arr := charIdxMap[target[idx]-'a']
        if len(arr) == 0 {
            return -1
        }
        n := len(arr)
        if currIdx > arr[n-1] {
            res++
            currIdx = -1
        }
        currIdx = binSearch(arr, currIdx)
    }
    return res 
}