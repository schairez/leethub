
// 852. Peak Index in a Mountain Array
// time: O(logn)
// space: O(1)
func peakIndexInMountainArray(arr []int) int {
    n := len(arr)
    lo, hi := 0, n-1 
    for lo + 1 < hi {
        mid := lo + (hi-lo) >> 1
        if arr[mid] > arr[mid+1] {
            hi = mid
        } else {
            lo = mid
        }
    }
    if arr[lo] > arr[lo+1] {
        return lo
    }
    return hi
}