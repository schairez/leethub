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
            //found = true
            if !found { found = true } 
            start := end -1
            for start >= 0 && s[start] != c {
                cand := abs(start - charIdx)
                fmt.Println(start, cand)
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


/*

    start, end := -1, 0


    for end < n {
        if s[end] == c {
            charIdx = end
            for {
                start++
                if start == n || s[start] == c {
                    break
                }
                fmt.Println("yo", start)
                res[start] = abs(start - charIdx)
                fmt.Println(res)
                //start++
            }
        } else {
            start++
            if start >= n || start == -1 {
                break
            }
            res[start] = abs(start - charIdx)
            start++
        }
        /*
        if start < n {
        } else {
            start = end+1
        }
        end++
    }
        */

