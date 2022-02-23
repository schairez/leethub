


// time: O(n^2) - at each recursive step we have O(n) build map op
// worst case we do N work for N recursive calls deep
// but on avg we perform better than the brute force, on avg more like O(nlogn) since
// shorter expected recursive depth
// space: O(n) recursive call stack + build map space; avg more like O(logn)

func max(a, b int) int { if a >= b { return a}; return b}

func longestSubstring(s string, k int) int {
    var helper func(left, right int) int
    helper = func(left, right int) int {
        if right - left < k {
            return 0
        }
        freq := [26]int{}
        for i := left; i < right; i++ {
            freq[s[i]-'a']++
        }
        for partLeft := left; partLeft < right; partLeft++ {
            if freq[s[partLeft]-'a'] < k {
                partRight := partLeft + 1
                for partRight < right && freq[s[partRight]-'a'] < k {
                    partRight++
                }
                return max(helper(left, partLeft), helper(partRight, right))
            }
        }
        return right - left
    }
    
    return helper(0, len(s))
}