/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     public var val: Int
 *     public var left: TreeNode?
 *     public var right: TreeNode?
 *     public init() { self.val = 0; self.left = nil; self.right = nil; }
 *     public init(_ val: Int) { self.val = val; self.left = nil; self.right = nil; }
 *     public init(_ val: Int, _ left: TreeNode?, _ right: TreeNode?) {
 *         self.val = val
 *         self.left = left
 *         self.right = right
 *     }
 * }
 */

class Solution {
    func upsideDownBinaryTree(_ root: TreeNode?) -> TreeNode? {
        if root == nil || root?.left == nil && root?.right == nil {
            return root
        }
        var currNode, prevLeft, prevRight: TreeNode?
        var prevRoot, nextLvlLeft, nextLvlRight: TreeNode? 
        currNode = root
        while true {
            nextLvlLeft = currNode?.left; nextLvlRight = currNode?.right
            currNode?.left = prevRight; currNode?.right = prevRoot 
            if nextLvlLeft == nil {
                break
            }
            prevRoot = currNode
            currNode = nextLvlLeft
            prevLeft = nextLvlLeft; prevRight = nextLvlRight
        }
        return currNode
        
    }
}