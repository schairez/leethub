/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
    delNodesMap := make(map[int]struct{},0)
    for _, val := range to_delete {
        delNodesMap[val] = struct{}{}
    }
    res := []*TreeNode{}
    if _, ok := delNodesMap[root.Val]; !ok {
        res = append(res, root)
    }
    preorder(root, delNodesMap,func(node *TreeNode) {
        if node == nil {
            return
        }
        if _, ok := delNodesMap[node.Val]; !ok {
            res = append(res, node)
        }
    })
    return res
    
}
//preorder - DLR
func preorder(node *TreeNode, delNodesMap map[int]struct{},
              addChildToResFn func(node *TreeNode)) {
    if node == nil {
        return
    }
    if _, ok := delNodesMap[node.Val]; ok {
        addChildToResFn(node.Left)
        preorder(node.Left, delNodesMap, addChildToResFn)
        node.Left = nil
        addChildToResFn(node.Right)
        preorder(node.Right, delNodesMap, addChildToResFn)
        node.Right = nil
    } else {
        if node.Left != nil {
            nextNode := node.Left
            if _, ok := delNodesMap[node.Left.Val]; ok {
                node.Left = nil
            }
            preorder(nextNode, delNodesMap,addChildToResFn)
        }
        if node.Right != nil {
            nextNode := node.Right
            if _, ok := delNodesMap[node.Right.Val]; ok {
                node.Right = nil
            }
            preorder(nextNode, delNodesMap,addChildToResFn)
        }
    }
}