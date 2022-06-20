// 792. Number of Matching Subsequences
// time: O(|s|) to append indices to mapping + O(n*wlog(s)) where n = lenWords; w = maxLenWord; s = len haystack
// thus time ≈ O(n*wlog(s))
// space: O(26*|s|) ≈ O(|s|)
func numMatchingSubseq(s string, words []string) int {
    var charIdxMap [26][]int
    for idx := range s {
        charIdxMap[s[idx]-'a'] = append(charIdxMap[s[idx]-'a'], idx)
    }
    res := 0
    for _, word := range words {
        currIdx := 0
        for wIdx := range word {
            arr := charIdxMap[word[wIdx]-'a']
            if n := len(arr); n == 0 || arr[n-1] < currIdx {
                currIdx = -1
                break
            }
            currIdx = binSearch(arr, currIdx)
            if currIdx == -1 {
                break
            }
            currIdx++
        }
        if currIdx != -1 {
            res++
        }
    }
    return res
}


// invariant: return first val >= t if exists
func binSearch(arr []int, target int) int {
    n := len(arr)
    if n == 0 {
        return -1
    }
    lo, hi := 0, n-1
    for lo+1 < hi {
        mid := lo + (hi-lo) >> 1
        if arr[mid] >= target {
            hi = mid
        } else {
            lo = mid
        }
    }
    if arr[lo] >= target {
        return arr[lo]
    } else if arr[hi] >= target {
        return arr[hi]
    }
    return -1
}