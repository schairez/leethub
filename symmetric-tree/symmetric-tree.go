//n, m total nodes in tree a and b ( copy of tree a)
//time: O(n + n) -> (O(2n)) -> O(n)
//space: O(n)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSymmetric(root *TreeNode) bool {
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
    