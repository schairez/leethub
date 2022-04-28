
func permuteUnique(nums []int) [][]int {
    n := len(nums)
    var (
        res [][]int
        dfs func(int)
    )
    
    dfs = func(idx int) {
        if idx == n {
            tmp := make([]int, n)
            copy(tmp, nums)
            res = append(res, tmp)
        } 
        
        set := make(map[int]struct{})
        for i := idx; i < n; i++ {
            if _, exists := set[nums[i]]; !exists {
                nums[i], nums[idx] = nums[idx], nums[i]
                dfs(idx+1)
                nums[i], nums[idx] = nums[idx], nums[i]
                set[nums[i]] = struct{}{}
            }
        }
    }
    
    dfs(0)
    return res
}