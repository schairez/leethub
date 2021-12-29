//time: O(n)
//space: O(1)

//dutch flag algo approach
//sort in-place
//precondition: int vals 0,1,2 only
func sortColors(nums []int)  {
    swap := func(i, j int) {
        nums[i], nums[j] = nums[j], nums[i]
    }
    n := len(nums)
    midVal := 1 //mid vals should be 1
    //two pointer approach
    i, j := 0, 0
    //topIdx captures vals = 2
    topIdx := n-1
    for j <= topIdx {
        switch {
        case nums[j] < midVal:
            swap(i, j)
            i++
            j++
        case nums[j] > midVal:
            swap(j, topIdx)
            topIdx--
        default: //if nums[j] == midVal -> idx j in right pos
            j++
        }
    }
    
}