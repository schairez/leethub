
// 387. First Unique Character in a String
// s consists of only lowercase English letters.
// time: O(2n) ≈ O(n)
// space: O(1)

func firstUniqChar(s string) int {
    n := len(s)
    var charMap [26]int
    i := 0
    for i < n {
        charMap[s[i]-'a']++
        i++
    }
    i = 0
    for i < n {
        if charMap[s[i]-'a'] == 1 {
            return i
        }
        i++
    } 
    return -1
}