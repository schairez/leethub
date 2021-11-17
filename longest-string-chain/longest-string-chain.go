import (
    "sort"
    "fmt"
)



//checks candidate word
//w2 is subset of w1
func isValidPredecessor(w1, w2 string) bool {
    fmt.Println(w1, w2)
    fmt.Println(len(w1), len(w2))
    i, j := 0, 0
    for j < len(w1) && i < len(w2) {
        if w1[j] == w2[i] {
            i++
        }
        j++
    }
    return i == len(w2)
}

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
        for j:=i-1; j>=0; j-- {
            if len(words[j]) + 1 < len(words[i]) {
                break
            } 
            if len(words[j]) == len(words[i]) {
                continue
            }
            if isValidPredecessor(words[i], words[j]) {
                fmt.Println("check")
                dp[i] = max(dp[i], dp[j]+1)
                res = max(res, dp[i])
            }
        }
    }
    fmt.Println(dp)
    return res
}