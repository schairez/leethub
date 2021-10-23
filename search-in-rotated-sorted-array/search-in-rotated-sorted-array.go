//time: O(logN)
//space: O(1)

func search(nums []int, target int) int {
    //n := len(nums)
    l := 0
    r := len(nums) -1
    for l <= r {
        m := l + (r-l) / 2 
        if nums[m] == target { return m }
        //if arr[m+1] < arr[m] && arr[m+1] >= target { //pivot 
        if nums[l] <= nums[m] {
            if target >= nums[l] && target < nums[m] {
                r = m-1
            } else {
                l = m+1
            }
        } else {
            if target > nums[m] && target <= nums[r] {
                l = m+1
            } else {
                r = m-1
            }
        }
        
    }
    return -1
    
}

