
//time: O(size of num)
//space:O(1)

/*
121 
121 % 10 -> 1 
121 / 10 -> 12 
12 % 10 -> 2 
12 / 10 -> 1 
xRev = 0 
xRev = xRev * 10 + (121 % 10 ) -> 1 
xRev = 1 *10 + (12 % 10) -> 12 
xRev = 12 *10 + 1 % 10 -> 121
*/


func isPalindrome(x int) bool {
    if x < 0 { return false}
    xRev := 0 
    for tmp := x; tmp != 0; tmp /= 10 {
        xRev = xRev *10 + (tmp % 10)
    }
    return xRev == x
    
}