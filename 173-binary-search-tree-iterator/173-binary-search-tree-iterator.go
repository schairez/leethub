/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
    nums []int
}


func Constructor(root *TreeNode) BSTIterator {
    nums := inorderMorris(root)
    return BSTIterator{nums}
}

// overkill morris traversal
func inorderMorris(root *TreeNode) []int {
    var (
        currNode *TreeNode
        predNode *TreeNode
        res      []int
    )
    currNode = root
    for currNode != nil {
        if currNode.Left != nil {
            predNode = currNode.Left
            for predNode.Right != nil && predNode.Right != currNode { 
                predNode = predNode.Right
            }
            if predNode.Right == nil {
                predNode.Right = currNode
                currNode = currNode.Left
            } else {
                predNode.Right = nil
                res = append(res, currNode.Val)
                currNode = currNode.Right
            }
        } else {
            res = append(res, currNode.Val)
            currNode = currNode.Right
        }
    }
    
    return res
}


func (this *BSTIterator) Next() int {
    v := this.nums[0]
    this.nums = this.nums[1:]
    return v
}


func (this *BSTIterator) HasNext() bool {
    return len(this.nums) != 0
}


/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */