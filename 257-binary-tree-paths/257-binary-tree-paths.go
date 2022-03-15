/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func binaryTreePaths(root *TreeNode) []string {
    var res []string
    dfs(root, []byte{}, &res)
    return res
}

/*
The number of nodes in the tree is in the range [1, 100].
*/

func dfs(node *TreeNode, path []byte, res *[]string) {
    path = strconv.AppendInt(path, int64(node.Val), 10)
    if node.Left == nil && node.Right == nil {
        tmp := make([]byte, len(path))
        copy(tmp, path)
        *res = append(*res, string(tmp))
        return
    }
    if node.Left != nil {
        dfs(node.Left, append(path, '-', '>'), res)
    }
    if node.Right != nil {
        dfs(node.Right, append(path, '-', '>'), res)
    }
}


