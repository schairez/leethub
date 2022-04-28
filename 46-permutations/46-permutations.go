func permute(nums []int) [][]int {
    n := len(nums)
    var (
        res [][]int
        dfs func(int)
    )
    
    dfs = func(idx int) {
        if idx == n-1 {
            res = append(res, append([]int{}, nums...))
            return
        }
        for i := idx; i < n; i++ {
            nums[i], nums[idx] = nums[idx], nums[i]
            dfs(idx+1)
            nums[i], nums[idx] = nums[idx], nums[i]
        }
    }
    dfs(0)
    
    return res
}