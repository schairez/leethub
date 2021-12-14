
/*

- a closingP cannot occur before an openingP

n = 2
left = 2
right = 0

                    f("", 2, 0)
                        |
                    l-- | r++
                   f("(", 1, 1)
                    /           \
             l-- | r++           |r--
           f("((",0,2)         f("(), 1, 0)
              /                      \
              |r--                 l--|r++
      f("(()",0,1)               f("()(",0,1) 
            |                         |
            |r--                      | r--
       f("(())",0,0)              f("()()",0,0)





*/
//backtrack approach
//runtime is bounded by n-th catalan number
//time: O(4^n/ sqrt(n))
//space: O(4^n/ sqrt(n))

func generateParenthesis(n int) []string {
    const (
        open byte = '('
        closed byte = ')'
    )
    //constraint: 1 <= n <= 8
    var res []string
    var dfs func(combo []byte, left, right int)
    //condition: l and r >= 0
    dfs = func(combo []byte, left, right int) {
        if left == 0 && right == 0 {
            tmp := make([]byte, len(combo))
            copy(tmp, combo)
            res = append(res, string(tmp))
            return
        }
        if left != 0 {
            dfs(append(combo, open), left-1, right+1)
        }
        if right != 0 {
            dfs(append(combo, closed), left, right-1)
        }
    }
    dfs([]byte{}, n, 0)
    
    return res
    
}




















