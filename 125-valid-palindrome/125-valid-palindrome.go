// 125. Valid Palindrome
// time: O(n)
// space: O(1)


func isPalindrome(s string) bool {
    alphabet := newAlphaNumTable()
    n := len(s)
    i, j := 0, n -1
    for i < j {
        alphaNum1, err := alphabet.alphaNumByte(s[i])
        if err != nil {
            i++
            continue
        } 
        alphaNum2, err := alphabet.alphaNumByte(s[j])
        if err != nil {
            j--
            continue
        } 
        if alphaNum1 != alphaNum2 {
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

// returns lower alpha char if valid
// else returns back numeric char is valid
// if not valid returns err
func (a alphaNumTable) alphaNumByte(b byte) (byte, error) {
    if v := int(b - 'a'); v >= 0 && v < 26 {
        return a.alphaTbl[v], nil
    }
    if v := int(b - 'A'); v >= 0 && v < 26 {
        return a.alphaTbl[v], nil
    }
    if idx := int(b - '0'); idx >= 0 && idx <= 9 {
        return a.numTbl[idx], nil
    } 
    
    return byte('0'), fmt.Errorf("invalid alpha")
}
