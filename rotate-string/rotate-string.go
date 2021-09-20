import "fmt"

/*
time: 
- number of rotations is equal to length of string input s
- each time we are doing a for loop to modify the byte slice index value by
  the one after it for an iter equal to the size of the string input
- therefore O(n*n) ~ O(n^2)
    
space: 
- for each iteration we are converting our string to a slice of bytes since 
  strings are immutable
- thus creating a new byte sequence. therefore O(N) space complexity
*/

/*
ex1   goal = "cdeab"
shifts on 
 s = "abcde"
     "bcdea"
     "cdeab" *** -> true
*/

/*
ex2; goal = "abced"
s = "abcde"

    "bcdea" = rotation; numRotations++ --> 1 
    "cdeab" = rotation; numRotations++ --> 2 
    "deabc" = rotation; numRotations++ --> 3
    "eabcd" = rotation; numRotations++ --> 4 
    "abcde" --> a "360" degree rotation; if after n rotations we don't reach goal
    
*/

func rotateString(s string, goal string) bool {
    n := len(s) 
    if n == 1 {
        return false
    }
    if len(goal) != n {
        return false
    }
    res := false
    currRotation := rotateLeftByOne(s) 
    //if after n rotations we don't reach our goal we return false 
    numRotations := 1 
    for numRotations != n {
        numRotations++   
        if currRotation == goal {
            res = true 
            break
        }
        currRotation = rotateLeftByOne(currRotation)
        fmt.Println(currRotation)
    }
    return res
    
    
}
//counter-clockwise rotation 
func rotateLeftByOne(s string) string {
    n := len(s)
    b := []byte(s)
    newLastVal := b[0] 
    for i := 0; i < n -1; i++ {
        b[i] = b[i+1] 

    }
    b[n-1] = newLastVal
    return string(b)
    
}

