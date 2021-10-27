//time: O(n)
//space:0(1)

func removeElement(nums []int, val int) int {
    //precondition: 0 <=n.len <= 100
    n := len(nums)
    if n == 0 { return 0 }
    res := 0
    for i := range nums {
        if nums[i] != val {
            if res < i { nums[res] = nums[i] }
            res++
        } 
    }
    return res
}