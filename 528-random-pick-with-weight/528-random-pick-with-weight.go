// 528. Random Pick with Weight

type Solution struct {
    totalWt int
    prefixWtSum []int
}

// time: O(n)
// space: O(n)
func Constructor(w []int) Solution {
    n := len(w)
    totalWt := w[0]
    prefixWtSum := w
    for i := 1; i < n; i++ {
        totalWt += prefixWtSum[i]
        prefixWtSum[i] += prefixWtSum[i-1]
    }
    return Solution{totalWt, prefixWtSum}
}


// time: O(logn)
// space: O(1)
func (this *Solution) PickIndex() int {
    target := rand.Intn(this.totalWt)+1
    lo, hi := 0, len(this.prefixWtSum)-1
    // low_bound()
    for lo + 1 < hi {
        mid := lo + (hi-lo) >> 1
        if this.prefixWtSum[mid] >= target {
            hi = mid
        } else {
            lo = mid
        }
    }
    if this.prefixWtSum[lo] >= target {
        return lo
    }
    return hi
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */