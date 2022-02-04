/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//import "strconv"

func max(a, b int) int {if a >= b { return a}; return b}

//constraint: numNodes in [1, 2^10]

func printTree(root *TreeNode) [][]string {
    height := getTreeHeight(root)
    numCols := 1 << (height + 1) - 1 
    res := make([][]string, height+1)
    for r := range res {
        res[r] = make([]string, numCols)
    }
    preorder(root, height, 0, (numCols-1) >> 1, func(val, r, c int) {
        fmt.Println(val, r, c)
        res[r][c] = strconv.Itoa(val)
    })
    return res
}

//DLR
func preorder(node *TreeNode, height, row, col int, addToResFn func(val, row, col int)) {
    if node == nil {
        return
    }
    addToResFn(node.Val, row, col)
    
    if node.Left != nil {
        preorder(node.Left, height, row+1, col - (1 << (height - row - 1)), addToResFn )
    }
    if node.Right != nil {
        preorder(node.Right, height, row+1, col + (1 << (height - row - 1)), addToResFn)
    }
}


//postorder approach
//LRD
//time: O(logn) if balanced tree; O(n) if skewed
//space: O(logn) if balanced tree; O(n) if skewed
func getTreeHeight(root *TreeNode) int {
    if root == nil {
        return -1
    }
    lHeight := getTreeHeight(root.Left)
    rHeight := getTreeHeight(root.Right) 
    
    return max(lHeight, rHeight) + 1
}
