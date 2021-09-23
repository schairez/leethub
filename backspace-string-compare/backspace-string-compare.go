/*
strings are immutable so turning it into a byteSequence requires O(n) extra space
time: O(n)
space: O(n)
*/
import "fmt"

func backspaceCompare(s string, t string) bool {
    afterBackspace := func(text string) string {
        stack := []byte{}
        bSeq := []byte(text)
        for i:=0; i <len(bSeq); i++ {
            if bSeq[i] == '#' { 
                if  len(stack) > 0 {
                    stack = stack[:len(stack)-1]
                }
                continue
            }
            stack = append(stack, bSeq[i])
        } 
        fmt.Println(string(stack))
        return string(stack)
            
    }
    sString, tString := afterBackspace(s), afterBackspace(t)
    return sString == tString
    
}