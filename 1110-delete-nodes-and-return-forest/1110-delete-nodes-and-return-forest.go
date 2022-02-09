/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
    if root == nil {
        return []*TreeNode{}
    }
    delNodesMap := make(map[int]struct{}, len(to_delete))
    for _, val := range to_delete {
        delNodesMap[val] = struct{}{}
    }
    res := make([]*TreeNode, 0)
    if _, delRootCond := delNodesMap[root.Val]; !delRootCond {
        res = append(res, root)
    } 
    preorder(root, delNodesMap, func(node *TreeNode) {
        res = append(res, node)
    })
    return res
}

func preorder(node *TreeNode, delNodesMap map[int]struct{}, appendResFn func(node *TreeNode) ) *TreeNode {
    if node == nil {
        return nil
    }
    nextLeftNode := node.Left
    nextRightNode := node.Right
    if _, delNodeCond := delNodesMap[node.Val]; delNodeCond {
        node = nil
        if nextLeftNode != nil {
            if _, delNextLeftCond := delNodesMap[nextLeftNode.Val]; !delNextLeftCond {
                appendResFn(nextLeftNode)
            }
        }
        if nextRightNode != nil {
            if _, delNextRightCond := delNodesMap[nextRightNode.Val]; !delNextRightCond {
                appendResFn(nextRightNode)
            }
        }
        preorder(nextLeftNode, delNodesMap, appendResFn)
        preorder(nextRightNode, delNodesMap, appendResFn)
        return nil
    }
    node.Left = preorder(nextLeftNode, delNodesMap, appendResFn)
    node.Right = preorder(nextRightNode, delNodesMap, appendResFn)
    return node
}


















/*

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
*/