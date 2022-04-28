func permute(nums []int) [][]int {
    n := len(nums)
    var (
        dfs func(int)
        res [][]int
    )
    dfs = func(idx int) {
        if idx == n {
            res = append(res, append([]int{}, nums...))
            return
        }
        for i := idx; i < n; i++ {
            nums[idx], nums[i] = nums[i], nums[idx]
            dfs(idx+1)
            nums[idx], nums[i] = nums[i], nums[idx]
        }
    }
    dfs(0)
    return res
}