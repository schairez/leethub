func isNumeric(b byte) bool {
    return b >= '0' && b <= '9'
}

func evalRPN(tokens []string) int {
    var stack []int
    for i := range tokens {
        if len(tokens[i]) > 1 || isNumeric(tokens[i][0]) { 
            num, _ := strconv.Atoi(tokens[i])
            stack = append(stack, num)
        } else {
            var res int
            n := len(stack)
            num2, num1 := stack[n-1], stack[n-2]
            stack = stack[:n-2]
            oper := tokens[i]
            switch oper {
            case "+":
                res = num1 + num2
            case "-":
                res = num1 - num2
            case "*":
                res = num1 * num2
            case "/":
                res = num1 / num2
            }
            stack = append(stack, res)
        }
    }
    return stack[len(stack)-1]
}