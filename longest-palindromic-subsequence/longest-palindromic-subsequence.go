
/*

for palindromic subsequence
- if i > j return 0
- if i == j, return 1; len at least 1 when 1 byte char val
- else if s[i] == s[j], return 2 + f(i+1, j-1)
  - ex: bab, s[0] == s[2], therefore 2 + f(1,1) -> 2 + 1 -> 3
-  else if s[i] != s[j]; return the max of the left and
   right subtree; 
  - max(f(i+1, j), f(i, j-1))
  - ex: "dada" -> 3 
        3/  \3
       dad  ada 
     1 /      \ 1
      a        d
     
     bab 2+1 =3
      /
      a 1 
*/
func max(a, b int) int { if a >= b { return a}; return b}

//2^n for naive backtrack approach

//dp topdown
// time: O(n^2)
// space: O(n^2 + n) ~ O(n^2)

//memoized top-down approach
func longestPalindromeSubseq(s string) int {
    //constraint: 1 <= s.length <= 1000
    n := len(s)
    if n == 1 { return 1}
    memo := make([][]int, n)
    for row := range memo {
        memo[row] = make([]int, n)
    }
    var dfs func(i, j int) int
    dfs = func(i, j int) int {
        if i > j { return 0 }
        if lps := memo[i][j]; lps > 0 {
            return lps
        }
        if i == j {
            memo[i][j] = 1
        } else if s[i] == s[j] {
            memo[i][j] = 2 + dfs(i+1,j-1)
        } else {
            lChild := dfs(i+1,j)
            rChild := dfs(i, j-1)
            memo[i][j] = max(lChild, rChild)
        }
        return memo[i][j]
    }
    i, j := 0, n-1
    return dfs(i, j)
    
}

