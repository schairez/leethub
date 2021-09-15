/*
time: O(n); len of string
space: O(n); stack can possibly be filled with all open bracket types
*/

func isValid(s string) bool {
    if len(s) == 1 {
        return false
    }
    charsMap := map[string]string{
        "}" : "{",
        "]" : "[",
        ")" : "(",
    }
    stack := []string{} 
    for _, v := range s {
        if openBracket, ok := charsMap[string(v)]; ok {
            if len(stack) == 0 {
                return false
            }
            lifoItem := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            if  openBracket != lifoItem {
                return false
            }
            continue
        }
        stack = append(stack, string(v))
       
    }
    if len(stack) > 0 {
        return false
    }
    return true 
    
}