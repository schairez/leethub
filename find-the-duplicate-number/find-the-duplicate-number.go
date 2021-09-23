/*
time: O(n)
space: O(1)
*/


/*
        0    1  2  3     4 
nums = [1,  -3, -4, -2, -2]

@iter0
nums[abs(nums[0])] -> nums[1] --> 3
3 > 0 --> nums[abs(nums[0])] = -3
@iter1
nums[abs(nums[1])] -> nums[3] --> 2 
2 > 0 --> nums[abs(nums[1])] = -2
@iter2 
nums[abs(nums[2])] ->  nums[4] --> 2
2 >0 --> nums[abs(nums[2])] = -2
@iter3
nums[abs(nums[3])] -> nums[2] --> 4
4 > 0 --> nums[abs(nums[3])] = -4
@iter4
nums[abs(nums[4])] -> nums[2] --> -4
-4 < 0 **check** so duplicate = abs(nums[4])


*/

func findDuplicate(nums []int) int {
    var duplicate int 
    for i:=0; i < len(nums); i++ {
        if nums[abs(nums[i])] < 0 {
            duplicate = abs(nums[i]) 
            break
        } 
        nums[abs(nums[i])] = -1 * nums[abs(nums[i])]
    }       
    return duplicate
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a 
}