// 438. Find All Anagrams in a String
// time: O(n)
// space: O(n)

func findAnagrams(s string, p string) []int {
    n := len(s)
    k := len(p)
    isAnagramFn := anagramChecker(p)
    var sDict [26]int
    start, end := 0, 0
    var res []int
    for end < n {
        sDict[s[end] - 'a']++
        if end - start + 1 == k {
            if isAnagramFn(sDict) {
                res = append(res, start)
            }
            sDict[s[start]-'a']--
            start++
        }
        end++
    } 
    return res
}


func anagramChecker(p string) func([26]int) bool {
    // s and p consist of lowercase English letters
    k := len(p)
    var pDict [26]int 
    for i := 0; i < k; i++ {
        pDict[p[i] - 'a']++
    } 
    return func(sDict [26]int) bool {
        for i := 0; i < 26; i++ {
            if sDict[i] != pDict[i] {
                return false
            }
        }
        return true
    }
}


