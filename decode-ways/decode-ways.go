/*
postorder traversal (LRD)
-left -> choose 1
-right -> choose 2 
- note: paranthesis notes res++ 

"11106"
               ""
             /   \       i:j
           1      11   s[0:2]
         /       /   \       2:4
       1        1    10    s[j:j+2]
      / \    x /     /   \ 
    1    10    0   (6)     _
 x /     /           
   0    (6)
   
 s = "12"
             ""
           /    \
         (1)     (12)
 s = "226"
               ""
             /    \ 
          2        22
         / \     /    \ x
       2  (26)  (6)     nil
      / \ x
    (6)  nil
       
         
postorder { f(n.left); f(n.right); n.val; }  
l -> s[i:i+1]
r -> s[i: i+2]
*/
// no leading zeroes

import "strconv"


func numDecodings(s string) int {
    if len(s) == 0 || s[0] == '0' {
        return 0
    }
    //res := 0
    //helper(s, &res, 0, 1)
    //memo := make([]int, 0, len(s) + 1)
    memo := make([]int, len(s) + 1)
    return helper(s,memo, 0, 1 )
    //return res
    
}


//dfs preorder-inspired
// recursion memo-ized
func helper(s string, memo []int,  i, j int) int {
    if !validByte(s[i:j]) {
        return 0
    }
    if j == len(s) {
        return 1 
        //*res = (*res) + 1 
    }
    if i == len(s) -1 {
        return 1 
    }
    if memo[i] != 0 {
        return memo[i]
    }
    res := helper(s, memo, j, j+1)
    if validByte(s[i:j+1]) {
        res += helper(s, memo, i, j+1 )
    }
    memo[i] = res 
    return res 
}



//only A-Z range
//func validByte(b []byte) bool {
func validByte(b string) bool {
    if b[0] == '0' { return false} //no leading 0's
    num, _ := strconv.Atoi(string(b)) //input contains only digits and zeros (no invalid in)
    if num >= 1 && num <= 26 {
        return true 
    } 
    return false
}

/*

// below is recursion; non-optimized
//dfs preorder-inspired
func helper(s string, res *int,  i, j int) {
    if !validByte(s[i:j]) {
        return
    }
    if j == len(s) {
        *res = (*res) + 1 
    }
    if j + 1 <= len(s) {
        helper(s, res, j, j+1) //left -> choose 1
        helper(s, res, i, j+1) //right -> choose 2
    }
}
*/
/*
'1' -> 'A'
b  -> 'A'
val -> '1'             b = 'A'
to encode -----> val = b - 'A' + 1 -> val = 1 (int type)
//val := (b - 'A') + 1
                 
                 val int -> 1 
to decode -----> 1 = b - 'A' + 1 ---> val +'A'-1 = b; b = 'A'
1 = b - 'A' + 1 
val - 1 = b - 'A'
b = val - 1 - 'A'
*/

