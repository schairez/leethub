
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//n, m total nodes in tree a and b ( copy of tree a)
//time: O(n + n) -> (O(2n)) -> O(n)
//space: O(n)
func isSymmetricR(root *TreeNode) bool {
    var mirror func(a, b *TreeNode) bool
    mirror = func(a, b *TreeNode) bool {
        switch {
            case a == nil && b == nil: 
                return true 
            case a != nil && b != nil:
                return (a.Val == b.Val && 
                    mirror(a.Left, b.Right) && 
                    mirror(a.Right, b.Left))
            default:
                return false
        }
    }
    return mirror(root, root)
}
    

type NodePair struct {
    a, b *TreeNode
}
func (np *NodePair) Destruct() (*TreeNode, *TreeNode) {
    //note in prod add validation code here
    return np.a, np.b
}


//iterative v with q
//time: O(n); space: O(n)
func isSymmetric(root *TreeNode) bool {
    queue := []*NodePair{}
    queue = append(queue, &NodePair{root,root})
    for len(queue) > 0 {
        a, b := queue[0].Destruct()
        queue = queue[1:]
        switch {
            case a == nil && b == nil:
                continue
            case a == nil || b == nil:
                return false
            case a.Val != b.Val:
                return false
            default:
                queue = append(queue, &NodePair{a.Left, b.Right})
                queue = append(queue, &NodePair{a.Right, b.Left})
        }
    }
    return true
}


