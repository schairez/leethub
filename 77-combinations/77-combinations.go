func combine(n int, k int) [][]int {
    var (
        res [][]int
        dfs func(int, []int)
    )
    dfs = func(idx int, combo []int) {
        if len(combo) == k {
            tmp := make([]int, k)
            copy(tmp, combo)
            res = append(res, tmp)
            return
        }
        for i := idx; i <= n; i++ {
            combo = append(combo, i)
            dfs(i+1, combo)
            combo = combo[:len(combo)-1]
        }
    }
    
    combo := make([]int, 0, k)
    dfs(1, combo)
    
    return res
}