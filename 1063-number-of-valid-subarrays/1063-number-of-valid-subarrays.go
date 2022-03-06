
/*
         [1,4,2]
         [0]
         res += 1    //  [1]
         @i=1
         res += 2   // [4],  [1,4]
         
         @i=2
          res +=2   //[2], [1,4,2]
         

mono increasing 

*/

func validSubarrays(nums []int) int {
    var res int
    var stack []int
    for i := range nums {
        for {
            if len(stack) == 0 ||
            nums[i] >= nums[stack[len(stack)-1]] {
                break
            }
            stack = stack[:len(stack)-1]
        }
        stack = append(stack, i)
        res += len(stack)
    }
    return res
}