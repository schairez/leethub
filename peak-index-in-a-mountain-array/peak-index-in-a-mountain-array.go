// arr = [3,4,5,1]
// m -> 0+4/2 -> 2
// 
// Input: arr = [24,69,100,99,79,78,67,36,26,19]
// m -> 5


func peakIndexInMountainArray(arr []int) int {
    lo, hi := 0, len(arr)
    for hi - lo > 0 {
        m := lo + (hi -lo)/2
        if arr[m] >= arr[m+1] {
            hi = m
        } else if arr[m] < arr[m+1] {
            lo = m+1
        }
    }
    return lo
}