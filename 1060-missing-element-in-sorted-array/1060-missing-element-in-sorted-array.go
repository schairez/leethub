
// 1060. Missing Element in Sorted Array
// time: O(logn)
// space: O(1)

func missingElement(nums []int, k int) int {
    n := len(nums)
    lo, hi := 0, n-1
    for lo + 1 < hi {
        mid := lo + (hi-lo) >> 1
        // range numElems based on vals
        numInRange := nums[mid] - nums[lo] + 1 
        // capture numElems based on idx
        numInArr := mid - lo + 1
        // how many missing positives in range?
        numMissing := numInRange - numInArr
        if numMissing < k {
            lo = mid
            k -= numMissing
        } else {
            hi = mid
        }
    }
    numMissing := nums[hi] - nums[lo] + 1 - (hi - lo + 1)
    if numMissing >= k {
        return nums[lo] + k
    }
    // k > numMissing
    return nums[hi] + k - numMissing
}

