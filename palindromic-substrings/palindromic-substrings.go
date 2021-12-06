
/*
Constraints:
1 <= s.length <= 1000
s consists of lowercase English letters.

Input: s = "aaa"
Output: 6
Explanation: Six palindromic strings: "a", "a", "a", "aa", "aa", "aaa".

*/

//time: O(n^2)
//space: O(1)

//TODO: learn manacher's algo

func countSubstrings(s string) int {
    n := len(s)
    //checks palindromic substr, if so increments cntr
    //expands around center
    cntPalindromes := func(l, r int) int {
        cnt := 0
        for l >= 0 && r < n && s[l] == s[r] {
            cnt++
            l--
            r++
        }
        return cnt
    }
    res := 0
    for i := 0; i < n; i++ {
        res += cntPalindromes(i, i)
        res += cntPalindromes(i, i+1)
    }
    return res
}