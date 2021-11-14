/**
 * Definition for a binary tree Node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//time: O(n)
//space: O(1)

func isCompleteTree(root *TreeNode) bool {
    q := []*TreeNode{}
    q = append(q, root)
    var curNode *TreeNode 
    peekHead := q[0]
    for peekHead != nil {
        curNode = q[0]
        q = q[1:]
        q = append(q, curNode.Left)
        q = append(q, curNode.Right)
        peekHead = q[0]
    }
    for len(q) > 0 && q[0] == nil {
            q = q[1:]
    }
    return len(q) == 0
}

/*

root = [1,2,3,4,5,6]
        |
       q = [1]
       curNode = q.pop() //1
       q = [2,3]
       
       q = [3,4,5]
       q = [4,5,6,nil]
       q = [nil]
       for len(q) > 0 and peek is nil { pop}
       return len(q) == 0

*/



