/*
time: O(N^((T/M)+1))
space: O(T/M)
*/


func combinationSum(candidates []int, target int) [][]int {
    var res [][]int
    if len(candidates) == 0 { return res }
    backtrack(candidates, 0, target, []int{}, &res)
    return res
    
}

func backtrack(candidates []int, idx int, remaining int, combo []int, res *[][]int) {
    if remaining < 0 { return }
    if remaining == 0 {
        tmp := make([]int, len(combo))
        copy(tmp, combo)
        *res = append(*res, tmp)
        return
    }
    for i := idx; i < len(candidates);i++ {
        combo = append(combo, candidates[i])
        backtrack(candidates, i, remaining-candidates[i], combo, res)
        combo = combo[:len(combo)-1]
    }
}

