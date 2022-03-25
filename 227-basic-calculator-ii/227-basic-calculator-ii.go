
func isNumeric(b byte) bool {
    if v := int(b - '0'); v >= 0 && v <= 9 {
        return true
    }
    return false
}

func isWhitespace(b byte) bool {
    return b == ' '
}


func calculate(s string) int {
    var (
        stack    []int
        topElem  int
        num      int
        operator byte 
        //currChar byte
    )
    operator = '+'
    n := len(s)
    i := 0
    for i < n {
    //for i := 0; i < n; i++ {
        //currChar = s[i]
        for i < n && isWhitespace(s[i]) {
            i++
        }
        if i < n && isNumeric(s[i]) {
            num = 0
            for i < n  && isNumeric(s[i]) {
                num = num * 10 + int(s[i] - '0')
                i++
            }
            i--
            switch operator {
            case '+':
                stack = append(stack, num)
            case '-':
                stack = append(stack, -num)
            case '*':
                topElem, stack = stack[len(stack)-1], stack[:len(stack)-1]
                stack = append(stack, topElem * num )
            case '/':
                topElem, stack = stack[len(stack)-1], stack[:len(stack)-1]
                stack = append(stack, topElem / num )
            }
        } else if i < n && !isWhitespace(s[i]) {
            operator = s[i]
        }
        i++
    }
    res := 0
    for len(stack) != 0 {
        topElem, stack = stack[len(stack)-1], stack[:len(stack)-1]
        res += topElem 
    }
    return res
}
    