
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
    /*
    if minIntNoLSB := minInt /10; x < minIntNoLSB {
        return false
    }
    if maxIntNoLSB := maxInt/10; x > maxIntNoLSB
    */
    return true 
}

func reverse(x int) int {
    reverseX := 0
    isNegative := false 
    if x < 0 {
        isNegative = true 
        x = -1 * x
    }
    for x > 0 {
        lsb := x % 10
        newRes := reverseX * 10 + lsb 
        if ((newRes-lsb)/10 != reverseX) {
            return 0
        }
        if !checkValidI32(newRes) {
            return 0
        }
        x = x / 10
        reverseX = newRes
    }
    if isNegative {
        return -1 * reverseX
    }
    return reverseX 
    
    
    
    
    
}