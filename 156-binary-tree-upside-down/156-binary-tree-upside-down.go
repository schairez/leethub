/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


//newRoot = prevLeft
//newRoot.Right = prevRoot
//newRoot.Left = prevRoot.Right


//recursive approach
//time: O(n)
//space: O(height); avg(logn); worst: O(n)
func upsideDownBinaryTree(root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }
    if root.Left == nil && root.Right == nil {
        return root
    }
    finalRoot := upsideDownBinaryTree(root.Left)
    nextRoot := root.Left
    nextRoot.Left, nextRoot.Right  = root.Right, root
    root.Left, root.Right = nil, nil
    return finalRoot
} 
    

//iterative approach
//time: O(n)
//space:O(1)

func upsideDownBinaryTreeIter(root *TreeNode) *TreeNode {
    if root == nil || root.Left == nil && root.Right == nil {
        return root
    } 
    var (
        currNode, prevLeft, prevRight *TreeNode
        prevRoot, nextLvlLeft, nextLvlRight *TreeNode
    )
    currNode = root
    for {
        nextLvlLeft, nextLvlRight = currNode.Left, currNode.Right
        currNode.Left, currNode.Right = prevRight, prevRoot
        if nextLvlLeft == nil {
            break
        }
        prevRoot = currNode
        prevLeft, prevRight = nextLvlLeft, nextLvlRight 
        currNode = prevLeft
        
    }
    return currNode
} 
    