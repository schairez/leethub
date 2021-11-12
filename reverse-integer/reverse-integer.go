
/*
time: O(n)
space: O(1)
*/

func checkValidI32(x int) bool {
    minInt32 := -1 * (1 <<31) 
    maxInt32 := (1 << 31) - 1
    if x < minInt32 || x > maxInt32 {
        return false
    }
    return true 
}

func reverse(x int) int {
    reverse_x := 0
    isNegative := false 
    if x < 0 {
        isNegative = true 
        x = -1 * x
    }
    for x > 0 {
        reverse_x = reverse_x * 10 + x % 10 
        x = x / 10
        if !checkValidI32(reverse_x) {
            return 0
        }
    }
    if isNegative {
        return -1 * reverse_x
    }
    return reverse_x 
    
    
    
    
    
}