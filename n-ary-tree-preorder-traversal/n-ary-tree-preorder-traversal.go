//time: O(N) N is number of nodes
//space: O(N) iterative stack

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */
//pre (DLR)

func preorder(root *Node) []int {
    var res []int
    if root == nil { return res }
    stack := []*Node{root}
    for len(stack) > 0 {
        node := stack[len(stack)-1]    
        res = append(res, node.Val)
        stack = stack[:len(stack)-1]
        childrenNodes := node.Children
        for i:=len(childrenNodes)-1; i >=0; i-- {
            stack = append(stack, childrenNodes[i])
        }
    }
    return res
    
}