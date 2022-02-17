
//ex
//nums = [-1,0,1,2,-1,-4]
//sortedV -> [-4,-1,-1,0,1,2]

//num1 + num2 + num3 = 0
// num1 + num2 = -num3
//  target = -num3

//time: O(nlogn) + O(n*n) ≈ O(n*n)
//space: O(logn) for go's internal sort
func threeSum(nums []int) [][]int {
    n := len(nums)
    if n < 3 {
        return [][]int{}
    }
    res := make([][]int, 0)
    sort.Ints(nums) //O(nlogn) time
    for i := 0; i < n; i++ {
        //no more valid complements to the right of pos numbers
        if nums[i] > 0 { 
            break
        }
        target := - nums[i]
        start, end := i + 1, n-1
        if i == 0 || nums[i] != nums[i-1] {
            for start < end {
                twoSum := nums[start] + nums[end]
                if twoSum == target {
                    res = append(res, []int{nums[i], nums[start], nums[end]})
                    for start != end && nums[start] == nums[start+1] {
                        start++
                    } 
                    for start != end && nums[end] == nums[end-1] {
                        end--
                    }
                    start++
                    end--
                } else if twoSum < target {
                    start++
                } else {
                    end--
                }
            }
        }
    } 
    return res
}