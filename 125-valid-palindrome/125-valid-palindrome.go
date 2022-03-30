// time: O(n)
// space: O(n)

import "strings"

func isPalindrome(s string) bool {
    alphabet := newAlphaNumTable()
    n := len(s)
    var sb strings.Builder
    sb.Grow(n)
    for i := 0; i < n; i++ {
        if alphaByte, err := alphabet.lowerAlphaIfValid(s[i]); err == nil {
            sb.WriteByte(alphaByte)
        } else {
            if idx := int(s[i] - '0'); idx >= 0 && idx <= 9 {
                sb.WriteByte(alphabet.numTbl[idx])
            } 
        }
    }
    str := sb.String()
    return isStrPalindrome(str)
}

func isStrPalindrome(s string) bool {
    n := len(s)
    if n == 0 {
        return true
    }
    i, j := 0, n -1 
    for i < j {
        if s[i] != s[j] {
            return false
        }
        i++
        j--
    }
    
    return true 
}


type alphaNumTable struct {
    alphaTbl string
    numTbl  string
}
func newAlphaNumTable() alphaNumTable {
    alphaTbl := make([]byte, 26)
    for i := range alphaTbl {
        alphaTbl[i] = 'a' + byte(i)
    }
    numTbl := make([]byte, 10)
    for i := range numTbl {
        numTbl[i] = '0' + byte(i)
    }
    return alphaNumTable{string(alphaTbl), string(numTbl)}
}


func (a alphaNumTable) lowerAlphaIfValid(b byte) (byte, error) {
    if v := int(b - 'a'); v >= 0 && v < 26 {
        return a.alphaTbl[v], nil
    }
    if v := int(b - 'A'); v >= 0 && v < 26 {
        return a.alphaTbl[v], nil
    }
    return byte('0'), fmt.Errorf("invalid alpha")
}

