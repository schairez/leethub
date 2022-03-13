// 46. Permutations
// space: O(n!)
// time: O(n!)

func permute(nums []int) [][]int {
    n := len(nums)
    var (
        res [][]int
        dfs func(int, []int)
    ) 
    dfs = func(sIdx int, perm []int) {
        if sIdx == n {
            //tmp := make([]int, 0, n)
            res = append(res, append([]int(nil), perm...))
            return
        }
        for i := sIdx; i < n; i++ {
            perm[i], perm[sIdx] = perm[sIdx], perm[i]
            dfs(sIdx+1, perm)
            perm[sIdx], perm[i] = perm[i], perm[sIdx] 
        }
    }
    dfs(0, nums)
    
    return res
}