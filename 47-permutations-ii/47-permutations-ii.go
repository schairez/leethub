// 47. Permutations II
// time: O(n* n!)
// space: O(n* n!)

func permuteUnique(nums []int) [][]int {
    n := len(nums)
    freqMap := make(map[int]int)
    for _, num := range nums {
        freqMap[num]++
    }
    var (
        res [][]int
        dfs func([]int)
    )
    dfs = func(perm []int) {
        if len(perm) == n {
            tmp := make([]int, n)
            copy(tmp, perm)
            res = append(res, tmp)
            return
        }
        for numKey, cntVal := range freqMap {
            if cntVal == 0 {
                continue
            }
            perm = append(perm, numKey)
            freqMap[numKey]--
            dfs(perm)
            freqMap[numKey]++
            perm = perm[:len(perm)-1]
        }
    }
    
    dfs([]int{})
    return res
}