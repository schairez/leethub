/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// iterator-morris version

type BSTIterator struct {
    succNode, currIterNode *TreeNode
}


func Constructor(root *TreeNode) BSTIterator {
    iterator := BSTIterator{succNode: root }
    iterator.currIterNode = iterator.findNext()
    return iterator
}

// uses morris under the hood :P
func (this *BSTIterator) findNext() *TreeNode {
    var currIterNode, predNode *TreeNode
    for this.succNode != nil {
        if this.succNode.Left != nil {
            predNode = this.succNode.Left
            for predNode.Right != nil && predNode.Right != this.succNode {
                predNode = predNode.Right
            }
            if predNode.Right == this.succNode {
                predNode.Right = nil
                currIterNode = this.succNode
                this.succNode = this.succNode.Right
                break
            } else {
                predNode.Right = this.succNode
                this.succNode = this.succNode.Left
            }
        } else {
            currIterNode = this.succNode
            this.succNode = this.succNode.Right
            break
        }
    }
    return currIterNode
}

func (this *BSTIterator) Next() int {
    // condition: next will be called when valid next
    // so if not needed but here to avoid pointer dereferencing
    /*
    if !this.HasNext()  {
        return -1 
    }
    */
    val := this.currIterNode.Val
    this.currIterNode = this.findNext()
    return val
}


func (this *BSTIterator) HasNext() bool {
    return this.currIterNode != nil
}





/*


// morris with storing all node values as a preprocess

type BSTIterator struct {
    nums []int
}


func Constructor(root *TreeNode) BSTIterator {
    nums := inorderMorris(root)
    return BSTIterator{nums}
}

// time: O(3n) ≈ O(n)
// space: O(n) to cache all the node values into an array
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

*/

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */