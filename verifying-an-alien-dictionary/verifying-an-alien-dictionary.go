/*
time: O(n^2)
space: O(1) since hashMap is always constant space 26 char bytes
*/

import "fmt"

func isAlienSorted(words []string, order string) bool {
    alienCodes := newAlienCodes(order)
    isSorted := true 
    lPtr, rPtr := 0, 1 
    for rPtr != len(words) && isSorted {
        wordL, wordR := words[lPtr], words[rPtr]
        if isLess := alienCodes.less(wordL,wordR); !isLess {
            fmt.Println(wordL, wordR)
            return false
        }
        lPtr++
        rPtr++
    }
    return isSorted 
    
    
}
type alienCodes struct {
    codes map[byte]int
}

func newAlienCodes(order string) alienCodes {
    lexMap := make(map[byte]int, 26)
    //representing the order of the byte in order string
    for i:=0; i < len(order); i++ {
        var b byte = order[i]
        lexMap[b] = i
    }
    alien := alienCodes{codes: lexMap}
    return alien
}

func (a alienCodes) less(left, right string) bool {
    isSorted := true 
    //we iter length of smallest of two inputs 
    lenL, lenR := len(left), len(right)
    smallestLen := lenL
    if lenL > lenR {
        smallestLen = lenR
    }
    for i:=0; i < smallestLen; i++ {
        byteL, byteR := left[i], right[i]
        codeL, codeR := a.codes[byteL], a.codes[byteR]
        if codeL == codeR {
            continue
        } else if codeL > codeR {
            return false
        } else {
            return true 
        }
        
    }
    //in the case that len(l) > len(r); any char in l > the blank identifier '∅'
    if lenL > lenR {
        isSorted = false
    }
    return isSorted
}