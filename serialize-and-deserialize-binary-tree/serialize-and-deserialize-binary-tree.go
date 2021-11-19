/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

import (
    "strings"
    "strconv"
)


type Codec struct {
    dummyV int
}

func Constructor() Codec {
    //-1000 <= Node.val <= 1000

    return Codec{-1001}
}

func (this *Codec) genDummyNode() *TreeNode { return &TreeNode{this.dummyV, nil, nil} }


// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
    if root == nil {
        return ""
    }
    sb := strings.Builder{}
    q := []*TreeNode{root}
    for len(q) > 0 {
        pollNode := q[0]
        q = q[1:]
        sb.WriteString(strconv.Itoa(pollNode.Val))
        sb.WriteByte('|')
        if !(pollNode.Val == this.dummyV) {
            lChild := this.serializeNode(pollNode.Left)
            q = append(q, lChild)
            rChild := this.serializeNode(pollNode.Right)
            q = append(q, rChild)
        } 
    }
    return sb.String()
}

func (this *Codec) serializeNode(node *TreeNode) *TreeNode {
    if node == nil {
        return this.genDummyNode()
    }
    return node
}
                       


// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {    
    if data == "" { return nil }
    
    var nodeVals []string
    nodeVals = strings.Split(data, "|")
    deque := &Deque{make([]*TreeNode, 0, len(nodeVals))}
    
    v, _ := strconv.Atoi(nodeVals[0])
    root := &TreeNode{v,nil,nil}
    
    deque.addLast(root)
    for i:=1; i<len(nodeVals)-1; i+=2 {
        pollNode := deque.popFirst()
        lChild, _ := strconv.Atoi(nodeVals[i])
        rChild, _ := strconv.Atoi(nodeVals[i+1])
        if lChild != this.dummyV {
            pollNode.Left = &TreeNode{lChild, nil, nil}
            deque.addLast(pollNode.Left)
        }
        if rChild != this.dummyV {
            pollNode.Right = &TreeNode{rChild, nil, nil}
            deque.addLast(pollNode.Right)
        }
    }
    return root
    
}
type Deque struct {
	vals []*TreeNode
}

func (d *Deque) addLast(node *TreeNode) {
	d.vals = append(d.vals, node)
}
                       
func (d *Deque) popFirst() *TreeNode {
    firstItem := d.vals[0]
	d.vals = d.vals[1:]
	return firstItem
}

func (d *Deque) popLast() *TreeNode {
    popItem := d.vals[len(d.vals)-1]
	d.vals = d.vals[:len(d.vals)-1]
    return popItem
}


/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */