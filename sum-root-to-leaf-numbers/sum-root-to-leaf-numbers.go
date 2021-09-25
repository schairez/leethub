//time:O(n); all nodes in tree
//space: O(n) stack space allocated for each node

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//NOTE: 0 <= Node.val <= 9

func sumNumbers(root *TreeNode) int {
    //memo := make(map[int]int)
    //helper(root, &res, memo) 
    
    return helper(root, 0) 
}

//root = [4,9,0,5,1]
// number 4->9->5 as 495
// number 4->9->1 as 491
// number 4->0 as 40



//preorder (DLR) --> 

func helper(node *TreeNode, currSum int) int  {
    if node == nil { return  0 }
    
    currSum = node.Val + (currSum) * 10 
    
    if node.Left == nil && node.Right == nil {
        return currSum
    } 
    return  helper(node.Left, currSum) + helper(node.Right, currSum)
}



//assumes no exp < 0 bc indexing memo 
func powMemo(memo map[int]int, base, exp int) int {
    powerVal, ok := memo[exp]
    if ok { return powerVal}
    
    if exp == 0 { return 1 }
    nextExp := exp -1 
    val := base * powMemo(memo, base, nextExp)
    memo[exp] = val 
    return val
}
/*
//10^2
// f(10,2) -> return 10 *f(10,1) * f(10,0) = 10*10*1 = 100
// f(10,1) -> return 10 * f(10,0) = 10 * 1
// f(10,0) -> 1
func pow(base, exp int) int {
    if exp == 0 {
        return 1
    }
    return base * pow(base, exp--)
}
*/
