/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
    InOrderArr []int
}

//time: O(n)
//space: O(n); stack for inorder traversal
func Constructor(root *TreeNode) BSTIterator {
    arr := []int{}
    stack := []*TreeNode{}
    for root != nil || len(stack) > 0 {
        for root != nil {
            stack = append(stack, root)
            root = root.Left
        }
        root = stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        arr = append(arr, root.Val)
        root = root.Right
    } 
    return BSTIterator{InOrderArr: arr}
}

//time:O(1); space: O(1)
func (this *BSTIterator) Next() int {
    nodeVal := this.InOrderArr[0]
    this.InOrderArr = this.InOrderArr[1:]
    
    return nodeVal
    
}

//time:O(1); space: O(1)
func (this *BSTIterator) HasNext() bool {
    if len(this.InOrderArr) == 0 {
        return false
    }
    return true
}


/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */