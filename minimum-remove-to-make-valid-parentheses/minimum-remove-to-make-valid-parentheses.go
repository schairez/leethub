import "bytes"

//time: O(n)
//space: O(n)

func minRemoveToMakeValid(s string) string {
    bytesArr := []byte(s)
    buf := &bytes.Buffer{}
    stack := []int{}
    bytesDelMap := make(map[int]bool) 
    
    for idx, ch := range bytesArr {
        switch ch {
            case ')':
                if len(stack) == 0 {
                    bytesDelMap[idx] = true
                } else {
                    stack = stack[:len(stack)-1]
                }
            case '(':
                stack = append(stack, idx)
            default:
                continue
        }
    }
    for i := range stack {
        bytesDelMap[stack[i]] = true
    }
    for i := range bytesArr {
        if !bytesDelMap[i] { buf.WriteByte(bytesArr[i]) } 
    }
    
    return buf.String()
    
}

/*
Input: s = "(a(b(c)d)"
Output: "a(b(c)d)"

*/



//Input: s = "lee(t(c)o)de)"
// out: -> "lee(t(c)o)de" or "lee(t(co)de)" 
//valid: "()", "(A)", "(AB)"

