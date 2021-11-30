

func combinationSum3(k int, n int) [][]int {
    var res = [][]int{}
    var backtrack func(sumSoFar, nextNumStart int, combo []int) 
    backtrack = func(sumSoFar, nextNumStart int, combo []int) {
        if sumSoFar == n && len(combo) == k {
            tmp := make([]int, len(combo))
            copy(tmp, combo)
            res = append(res, tmp)
            return
        } else if sumSoFar > n {
            return
        }
        for i:=nextNumStart; i<9; i++ {
            sumSoFar+=i+1
            if sumSoFar > n {
                return
            }
            combo = append(combo, i+1)
            backtrack(sumSoFar, i+1, combo)
            combo = combo[:len(combo)-1]
            sumSoFar-=i+1
        }
    }
    backtrack(0, 0, []int{})
    
    return res
}