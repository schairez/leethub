// 90. Subsets II
// time: O(n* 2^n) 
// space: O(n* 2^n) + O(logn) to sort ≈ O(n*2^n) 

func subsetsWithDup(nums []int) [][]int {
    sort.Ints(nums) // nlogn time with logn space
    n := len(nums)
    var (
        res [][]int
        dfs func(int)
    )
    subset := make([]int, 0, n)
    dfs = func(idx int) {
        res = append(res, append([]int{}, subset...))
        if idx == n {
            return
        }
        for i := idx; i < n; i++ {
            if i != idx && nums[i] == nums[i-1] {
                continue
            }
            subset = append(subset, nums[i])
            dfs(i+1)
            subset = subset[:len(subset)-1]
        }
    }
    dfs(0)
    return res
}