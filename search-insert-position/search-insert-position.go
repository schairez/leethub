//time : O(logN)
//space: O(1)

func searchInsert(nums []int, target int) int {
    binSearch := func(target int) int {
        lo, hi := 0, len(nums)
        for lo < hi {
            mid := lo + (hi-lo) >> 1
            switch {
            case target > nums[mid]:
                lo = mid+1
            case target <= nums[mid]:
                hi = mid
            }
        }
        return lo
    }
    return binSearch(target)
}

/*
m -> 0 + 3/2 -> 1


              F, T,T,T
Input: nums = [1,3,5,6], target = 2
Output: 1
               F,F,F,T
Input: nums = [1,3,5,6], target = 7
Output: 4





*/