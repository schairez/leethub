// time: O(n)
// space: O(1) b/c of tail recursion

/*
 i.                   j
["H","a","n","n","a","h"]

*/

func reverseString(s []byte)  {
    var helper func(s []byte, i, j int)
    helper = func(s []byte, i, j int) {
        if i > j {
            return
        }
        s[i], s[j] = s[j], s[i] 
        i++
        j--
        helper(s, i, j)
    }
    i, j := 0, len(s)-1
    helper(s, i, j)
    
}