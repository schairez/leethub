//time: O(nlogn + 2^n) ~ O(2^n) 
//space: O(n)

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
    var res [][]int
    sort.Ints(candidates) //O(nlogn)
    var backtrack func(sumSoFar, nextIdx int,combo []int)
    backtrack = func(sumSoFar, nextIdx int, combo []int) {
        if sumSoFar == target {
            tmp := make([]int, len(combo))
            copy(tmp, combo)
            res = append(res, tmp)
            return
        }
        for i:=nextIdx; i<len(candidates); i++ {
            if i > nextIdx && candidates[i] == candidates[i-1] {
                continue
            }
            if candidates[i] > target {
                break
            }
            sumSoFar += candidates[i]
            if sumSoFar > target {
                break
            }
           
            combo = append(combo, candidates[i])
            backtrack(sumSoFar, i+1, combo)
            combo = combo[:len(combo)-1]
            sumSoFar -=candidates[i]
        }
        
    }
    backtrack(0,0, []int{})
    return res
}
