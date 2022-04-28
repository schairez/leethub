
// 78. Subsets
// time: O(n*2^n)
// space: O(n*2^n)

func subsets(nums []int) [][]int {
    n := len(nums)
    // 2^n subsets
    res := make([][]int, 0, n << 1)
    subset := make([]int, 0, n)
    var dfs func(int)
    dfs = func(idx int) {
        res = append(res, append([]int{}, subset...))
        if idx == n {
            return
        }
        for i := idx; i < n; i++ {
            subset = append(subset, nums[i])
            dfs(i+1)
            subset = subset[:len(subset)-1]
        }
    }
    dfs(0)
    return res
}