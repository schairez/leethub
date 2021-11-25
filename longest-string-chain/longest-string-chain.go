import (
    "sort"
    "fmt"
)

/*
time: sorting O(nlogn); two for loops for the dp bottom up subproblems O(L^2) where L=len max word

*/



type ByLen []string
 
func (a ByLen) Len() int {
   return len(a)
}
 
func (a ByLen) Less(i, j int) bool {
   return len(a[i]) < len(a[j])
}
 
func (a ByLen) Swap(i, j int) {
   a[i], a[j] = a[j], a[i]
}
/*

//checks candidate word
//predecessorCandidate is subset of successorWord

true if predecessorCandidate is a predecessor of successorWord

*/

func isValidPredecessor(successorWord, predecessorCandidate string) bool {
    cntMatchPredecessor, successorCharIdx := 0, 0
    for successorCharIdx < len(successorWord) && cntMatchPredecessor < len(predecessorCandidate) {
        if successorWord[successorCharIdx] == predecessorCandidate[cntMatchPredecessor] {
            cntMatchPredecessor++
        }
        successorCharIdx++
    }
    return cntMatchPredecessor == len(predecessorCandidate)
}


func longestStrChain(words []string) int {
    max := func(a, b int) int {
        if a >= b { return a }
        return b
    }
    res := 1
    n := len(words)
    sort.Sort(ByLen(words))
    dp := make([]int, len(words))
    for i := range dp {
        dp[i] = 1
    }
    for i:=1; i < n; i++ {
        successor := words[i]
        for j:=i-1; j>=0; j-- {
            predecessorCand := words[j]
            //only consider words with one less char in len
            if len(predecessorCand) + 1 < len(successor) {
                break
            } 
            if len(predecessorCand) == len(successor) {
                continue
            }
            if isValidPredecessor(successor, predecessorCand) {
                dp[i] = max(dp[i], dp[j]+1)
                res = max(res, dp[i])
            }
        }
    }
    fmt.Println(dp)
    return res
}