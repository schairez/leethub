/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

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
    

//newRoot = prevLeft
//newRoot.Right = prevRoot
//newRoot.Left = prevRoot.Right

//iterative approach
//time: O(n)
//space:O(1)
/*
func upsideDownBinaryTree(root *TreeNode) *TreeNode {
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
    
 */   
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
/*    
    
func upsideDownBinaryTree(root *TreeNode) *TreeNode {
    //newRoot -> prevLeft
    //newRoot.Right -> prevRoot
    //newRoot.Left = prevRight
    var (
        prevRoot, prevRight *TreeNode
        currNode, nextLvlLeft, nextLvlRight *TreeNode
    )
    currNode = root
    //continue traversing until next level children is nil
    for {
        //cache the child left and right nodes for next iter
        //these will become the next root and nex
        nextLvlLeft, nextLvlRight = currNode.Left, currNode.Right
        currNode.Left, currNode.Right = prevRight, prevRoot
        if nextLvlLeft == nil {
            break
        }
        currNode, prevRoot, prevRight  = nextLvlLeft, currNode, nextLvlRight
    }
    return currNode
}
    

func upsideDownBinaryTreeRecursive(root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }
    if root.Left == nil && root.Right == nil {
        return root
    }
    finalRootNode := upsideDownBinaryTree(root.Left)
    newRootNode := root.Left
    newRootNode.Right = root
    newRootNode.Left = root.Right
    root.Left, root.Right = nil, nil
    
    return finalRootNode
}
*/