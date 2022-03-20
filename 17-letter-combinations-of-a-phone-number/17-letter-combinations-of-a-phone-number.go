

func letterCombinations(digits string) []string {
    if len(digits) == 0 {
        return []string{}
    }
    letterMap := [10]string{
        "",
        "",
        "abc",
        "def",
        "ghi",
        "jkl",
        "mno",
        "pqrs",
        "tuv",
        "wxyz",
    }
    n := len(digits)
    var res []string
    var dfs func(start int, soFar []byte)
    
    dfs = func(start int, soFar []byte) {
        if len(soFar) == n {
            tmp := make([]byte, n)
            copy(tmp, soFar)
            res = append(res, string(tmp))
            return
        }
        //
        for i := start; i < n; i++ {
            idx, _ := strconv.Atoi(string(digits[i]))
            for j := 0; j < len(letterMap[idx]); j++ {
                v := letterMap[idx][j]
                soFar = append(soFar, v)
                dfs(i+1, soFar) //
                soFar = soFar[:len(soFar)-1]
            }
        }
    } 
    dfs(0, []byte{})
    
    
    return res
    
}