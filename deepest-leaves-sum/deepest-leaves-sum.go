/*
time: O(N)
space: O(N)

*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
traverse down level by level bfs
keep track of start idx of each level


root = [1,2,3,4,5,null,6,7,null,null,null,null,8]
q = [1]
1 [2,3]
*/

func deepestLeavesSum(root *TreeNode) int {
    queue := []*TreeNode{}
    //level := 0
    numNodesNextDepth := 0
    numNodesCurrDepth := 1 
    sumNodesCurrDepth := 0 
    node := root
    queue = append(queue, node)
    for len(queue) > 0 {
        node = queue[0]
        queue = queue[1:]
        if node.Left != nil {
            queue = append(queue, node.Left)
            numNodesNextDepth++
        }
        if node.Right != nil {
            queue = append(queue, node.Right)
            numNodesNextDepth++
        }
        numNodesCurrDepth--
        sumNodesCurrDepth += node.Val
        
        if numNodesCurrDepth == 0 {
            numNodesCurrDepth = numNodesNextDepth
            numNodesNextDepth = 0 
            if len(queue) > 0 {
                sumNodesCurrDepth = 0 
            }
            
        }
    
    }
    return sumNodesCurrDepth
    
}