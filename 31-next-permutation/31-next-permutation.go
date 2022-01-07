//time: O(n)
//space: O(1)

func nextPermutation(nums []int)  {
    n := len(nums)
    i := n-2
    for i >= 0 && nums[i] >= nums[i+1] {
        i--
    }
    j := n-1
    if i >= 0 {
        for nums[i] >= nums[j] && j >= 0 {
            j--
        }
        nums[j], nums[i] = nums[i], nums[j]
    }
    lPtr, rPtr := i+1, n-1 
    //reverse slice in place
    for lPtr < rPtr {
        nums[lPtr], nums[rPtr] = nums[rPtr], nums[lPtr]
        lPtr++
        rPtr--
    }
}

/*

[3,2,1]
   | |
   i j
   swap if 
 
 [1,2,3] to [1,3,2]
    | |
    i j  while nums[i] > nums[i+1] i--
    
    
    [1,1,5]    []
       | |
       i j
       
       
       
*/


























