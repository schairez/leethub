// 20. Valid Parentheses
// time: O(n)
// space: O(n)

func isValid(s string) bool {
    var stack []byte
    for i := range s {
        switch ch := s[i]; ch {
        case '(':
            stack = append(stack, ')')
        case '[':
            stack = append(stack, ']')
        case '{':
            stack = append(stack, '}')
        case '}', ']', ')':
            if len(stack) == 0 {
                return false
            }
            if v := stack[len(stack)-1]; v != ch {
                return false
            }
            stack = stack[:len(stack)-1]
        }
    }
    return len(stack) == 0
}