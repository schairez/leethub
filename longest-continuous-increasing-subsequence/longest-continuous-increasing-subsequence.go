func max(a, b int) int { if a > b { return a}; return b}

//time : O(n)
//space: O(1)
//bottom up with no dp table
func findLengthOfLCIS(nums []int) int {
    n := len(nums)
    res := 1
    cnt :=1
    for i:=1; i < n; i++ {
        if nums[i] > nums[i-1] {
            cnt++
        } else {
            cnt = 1
        }
        res = max(res, cnt)
    }
    return res
} 
//time : O(n)
//space: O(n)

func findLengthOfLCISTab(nums []int) int {
    //1 <= nums.length <= 104
    n := len(nums)
    res := 1
    
    cnt := 1 //cnt continuous increasing
    for i:=1; i < n; i++ {
        if nums[i] > nums[i-1] {
            cnt++
        } else {
            cnt = 1
        }
        res = max(cnt, res)
    }
   return res 
}

/*
Input: nums = [1,3,5,4,7]


*/
