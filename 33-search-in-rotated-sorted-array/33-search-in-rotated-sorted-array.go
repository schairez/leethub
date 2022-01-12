/*
Input: nums = [4,5,6,7,0,1,2], target = 0
Output: 4

*/
//time: O(logn)
//space: O(1)
//note: if input arr is rotated k times and 1 <= k < len(nums)
// therefore, initially, we have nums[0] > nums[n-1]
//ex: [0,1,2,4,5,6,7]
//rotated at k =1; [1,2,4,5,6,7,0]
//rotated at k=2; [2,4,5,6,7,0,1]
//...
// ... k = 4; [5,6,7,0,1,2]
// k = n-1 = 6; [2,4,5,6,7,0,1]
//all have nums[loIdx] > nums[hiIdx]

//nums may or may not be pivoted 
func search(nums []int, target int) int {
    lowIdx, highIdx := 0, len(nums)-1
    midIdx := 0
    for lowIdx != highIdx {
        midIdx = lowIdx + (highIdx-lowIdx) >> 1
        //nums may or may not be pivoted 
        if nums[midIdx] >= nums[lowIdx] {
            if target <= nums[midIdx] && target >= nums[lowIdx] {
                highIdx = midIdx
            } else {
                lowIdx = midIdx+1
            }
        } else {
            if target > nums[midIdx] && target <= nums[highIdx] {
                lowIdx = midIdx+1
            } else {
                highIdx = midIdx
            }
        }
    }
    if nums[lowIdx] == target {
        return lowIdx
    }
    return -1
}


/*

10/23/2021 version

//time: O(logN)
//space: O(1)
//tmpl#1
func search(nums []int, target int) int {
    l := 0
    r := len(nums) -1
    for l <= r {
        m := l + (r-l) / 2 
        if nums[m] == target { return m }
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
*/
