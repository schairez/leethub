/*
time: O(nlogn) bc of sort
space: O(1) in place mutable
*/

import (
    "sort"
    "strings"
)

func isNumericByte(b byte) bool {
    if b >= byte('0') && b <= byte('9') { 
        return true
    }
    return false
}

func reorderLogFiles(logs []string) []string {
    //stable sort to preserve initial order
    sort.SliceStable(logs, func(i, j int) bool {
        var isW1Numeric, isW2Numeric bool 
        w1 := logs[i]
        w2 := logs[j]
        w1s := strings.SplitAfterN(w1, " ", 2)
        w2s := strings.SplitAfterN(w2, " ", 2)
        ident1, content1 := w1s[0], w1s[1]
        ident2, content2 := w2s[0], w2s[1]
        //digit-logs all content is numeric
        isW1Numeric = isNumericByte(content1[0])
        isW2Numeric = isNumericByte(content2[0])
        var retVal bool
        switch {
            case !isW1Numeric && !isW2Numeric:
                if content1 == content2 { 
                    retVal = ident1 < ident2
                    break
                }
                retVal = content1 < content2
            case !isW1Numeric && isW2Numeric:
                retVal = true //i is word so true, but j is num
            case !isW2Numeric && isW1Numeric:
                retVal = false //i is num, but j is word so j before i
            case isW1Numeric && isW2Numeric:
                retVal = i < j
                
            }
        return retVal        
        
    })
    return logs    
}