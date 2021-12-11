
/*

s = "abccccdd"
"a" ; res =1

"ab"; 1
"abc"; 1
"abcc"; 3 "cac"
"abccc"; 3 "ccc" or "cac"
"abcccc"; 5; "ccacc"
"abccccd"; 5; "ccacc"
"abccccdd"; 7; "dccaccd"
///////////////////////
//"aaabb" palindrome = "baaab"
//"deeeffe" palindrome = "feedeef" unival = 1; if < 2 return n
//"azfeeeee" largest charCntFreq = e @ freq=5; 3 uni vals {a,z,f}
//oddValsCnt cnt = 4; 
// n -(3 -1) = 8 -(4-1) = 5; eezee

*/

//time: O(n)
//space:O(n)

func max(a, b int) int { if a >= b { return a}; return b }

func longestPalindrome(s string) int {
    //assumes ascii cd pts for elem in [AZaz]
    getCodePtIdx := func(ch byte) int {
        if  ch >= 'A' && ch <= 'Z' {
            return int(ch - 'A') + 26
        }
        return int(ch - 'a')
    }
    n := len(s)
    //1 <= s.length <= 2000
    var charsCntMap [52]int
    for i :=0; i < len(s); i++ {
        charsCntMap[getCodePtIdx(s[i])]++
    }
    oddValsCnt := 0
    for _, cnt := range charsCntMap {
        if cnt % 2 == 1 {
            oddValsCnt++
        }
    }
    if oddValsCnt < 2 {
        return n
    }
    return n - (oddValsCnt-1)
    
}