
func decodeString(s string) string {
    res, _ := dfs(s, 0)
    return res
}

func dfs(s string, idx int) (string, int) {
    n := len(s)
    var sb strings.Builder
    numTimes := 0
    for idx < n {
        if isAlpha(s[idx]) {
            for idx < n && isAlpha(s[idx]) {
                sb.WriteByte(s[idx])
                idx++
            }
        } else if isDigit(s[idx]) {
            for idx < n && isDigit(s[idx]) {
                numTimes = numTimes*10 + int(s[idx]-'0')
                idx++
            }
        } else if s[idx] == '[' {
            nextStr, nextIdx := dfs(s, idx+1)
            for i := 0; i < numTimes; i++ {
                sb.WriteString(nextStr)
            }
            numTimes = 0
            idx = nextIdx
            //idx = idx + len(nextStr) + 2
        } else if s[idx] == ']' {
            return sb.String(), idx+1
        }
    }
    return sb.String(), idx+1
}

func isAlpha(b byte) bool {
    return b >= 'a' && b <= 'z'
}

func isDigit(b byte) bool {
    return b >= '0' && b <= '9'
}

