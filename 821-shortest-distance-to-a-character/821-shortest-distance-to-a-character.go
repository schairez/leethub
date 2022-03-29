

// 821. Shortest Distance to a Character
// time: O(n)
// space: O(1) if not counting res; O(n) if counting res as growth rate based on input
// but no aux D.S. were used ;)

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func shortestToChar(s string, c byte) []int {
    n := len(s)
    res := make([]int, n)
    var charIdx int
    found := false
    end := 0
    for end < n {
        if s[end] != c {
            if found {
                res[end] = abs(end - charIdx)
            }
        } else {
            charIdx = end
            if !found { found = true } 
            start := end -1
            for start >= 0 && s[start] != c {
                cand := abs(start - charIdx)
                if res[start] != 0 && res[start] < cand {
                    break
                }
                res[start] = cand
                start--
            }
        }
        end++
    }
    fmt.Println(end)
    
    return res
    
}
