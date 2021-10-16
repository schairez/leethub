//iterative
// time: O(len(t))
// space: O(1)
func isSubsequence(s string, t string) bool {
    //if subsequence; num of common valid chars == len(s) 
    lenS := len(s)
    lenT := len(t)
    lenCommon := 0
    for i:=0; i < lenT && lenCommon < lenS; i++ {
        if t[i] == s[lenCommon] {
            lenCommon++
        }     
    }
    return lenS == lenCommon
    
}
/*
//time: O(len(t)) //worst-case we iterate through all bytes of larger string
//space: O(n) //stack space

//is s subsequence of t; recursive
func isSubsequence(s string, t string) bool {
    if len(s) == 0 { return true }
    if len(t) == 0 { return false }
    
    if s[0] == t[0] {
        return isSubsequence(s[1:], t[1:])
    }
    return isSubsequence(s, t[1:])
    
}

*/