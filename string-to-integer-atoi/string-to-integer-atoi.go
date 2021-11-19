
func myAtoi(s string) int {
    sBytes := []byte(s)
    maxInt32 := (1 << 31) - 1
    minInt32 := -1 << 31
    //delete leading whitespace
    whitespace := byte(' ')
    //tabDelim := '\t'
    curIdx :=0
    for curIdx < len(sBytes) && sBytes[curIdx] == whitespace {
        curIdx++
    }
    if curIdx == len(sBytes) { return 0 } 
    
    signBit := 1
    if sBytes[curIdx] == '-' {
        signBit = -1
        curIdx++
    } else if sBytes[curIdx] == '+' {
        curIdx++
    } 
    res := 0
    for curIdx < len(sBytes) && sBytes[curIdx] >= '0' && sBytes[curIdx] <= '9' {
        numBit := int(sBytes[curIdx] - '0') 
        res = res*10 + numBit 
        if res * signBit > maxInt32 {
            return maxInt32
        } else if res * signBit < minInt32 {
            return minInt32
        } 
        curIdx++
    }
    return res*signBit
}

