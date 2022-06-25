
// 715. Range Module
// treap probabilistically BBST approach
// more implementations: https://github.com/schairez/leethub

// queryRange
// time: O(h) ≈ O(logn) expected time O(n) worst case but rare outcome
// space: O(h) ≈ O(logn) expected time O(n) worst case but rare outcome


// removeRange
// worst case removeRange inserts 2 childNodes on each split, but since BBST exp time: O(2logn) ≈ O(logn) 
// space: O(logn) exp but low probabilistic case is O(n)

// insertRange
// time: O(logn) exp but low probabilistic case is O(n) for both time and space
// space: O(logn) 


import "math/rand"

func init() {
    rand.Seed(time.Now().UnixNano())
}

// min heap property along with BST property
type TreapNode struct {
    HeapKey int
    Interval [2]int
    Child [2]*TreapNode
}

type RangeModule struct {
    root *TreapNode
}


func Constructor() RangeModule {
    return RangeModule{}
    
}


func (this *RangeModule) AddRange(left int, right int)  {
    this.root = insertNode(this.root, left, right)
    
}

func insertNode(node *TreapNode, left, right int) *TreapNode {
    if left >= right {
        return node
    }
    if node == nil {
        newNode := &TreapNode{Interval: [2]int{left, right}, HeapKey: rand.Intn(1 << 30) % 6662333 }
        return newNode
    }
    // disj cond 
    if right <= node.Interval[0] {
        node.Child[0] = insertNode(node.Child[0], left, right)
    } else if left >= node.Interval[1] {
        node.Child[1] = insertNode(node.Child[1], left, right)
    } else { // intersects
        node.Child[0] = insertNode(node.Child[0], left, node.Interval[0])
        node.Child[1] = insertNode(node.Child[1], node.Interval[1], right)
    }
    // check heap property and rotateUp accordingly
    for dir := 0; dir <= 1; dir++ {
        if node.Child[dir] != nil && node.Child[dir].HeapKey < node.HeapKey {
            node = rotateUp(node, dir)
        }
    }
    return node
}



func (this *RangeModule) QueryRange(left int, right int) bool {
    return treapSearch(this.root, left, right)
}


// check if treap covers the interval query
func treapSearch(node *TreapNode, left, right int) bool {
    if left >= right {
        return true
    }
    if node == nil {
        return false
    }
    // disj from currNode
    if left >= node.Interval[1] {
        return treapSearch(node.Child[1], left, right)
    } else if right <= node.Interval[0] {
        return treapSearch(node.Child[0], left, right)
    }
    // overlaps currNode
    if left >= node.Interval[0] && right <= node.Interval[1] {
        return true
    }
    // overlaps child nodes
    return treapSearch(node.Child[0], left, node.Interval[0]) && treapSearch(node.Child[1], node.Interval[1], right)
}


func (this *RangeModule) RemoveRange(left int, right int)  {
    this.root = removeRange(this.root, left, right)
}


func removeRange(node *TreapNode, left, right int) *TreapNode {
    if left >= right || node == nil {
        return node
    }
    // interval disj from currNode
    if node.Interval[1] <= left {
        node.Child[1] = removeRange(node.Child[1], left, right)
    } else if node.Interval[0] >= right {
        node.Child[0] = removeRange(node.Child[0], left, right)
    } else { //overlaps currNode
        node.Child[0] = removeRange(node.Child[0], left, node.Interval[0])
        node.Child[1] = removeRange(node.Child[1], node.Interval[1], right)
        node.Child[0] = insertNode(node.Child[0], node.Interval[0], left)
        node.Child[1] = insertNode(node.Child[1], right, node.Interval[1])
        node.Interval[0], node.Interval[1] = left, right 
        node = removeNode(node, node.Interval[0], node.Interval[1])
    }
    return node
}

func rotateUp(node *TreapNode, dir int) *TreapNode {
    childNode := node.Child[dir]
    node.Child[dir] = childNode.Child[dir ^ 1] 
    childNode.Child[dir ^ 1] = node
    return childNode 
}


func removeNode(node *TreapNode, left, right int) *TreapNode {
    if node == nil {
        return node
    }
    if node.Interval[0] == left && node.Interval[1] == right {
        lChild, rChild := node.Child[0], node.Child[1]
        // single parent cond
        if lChild == nil || rChild == nil {
            var retNode *TreapNode 
            if lChild == nil {
                retNode = rChild
                node.Child[1] = nil
            } else {
                retNode = lChild
                node.Child[0] = nil
            }
            node = nil
            return retNode
        }
        dir := 0
        if rChild.HeapKey < lChild.HeapKey {
            dir = 1
        }
        node = rotateUp(node, dir)
    }
    if node.Interval[0] >= right {
        node.Child[0] = removeNode(node.Child[0], left, right)
    } else {
        node.Child[1] = removeNode(node.Child[1], left, right)
    }
    return node
}



/**
 * Your RangeModule object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddRange(left,right);
 * param_2 := obj.QueryRange(left,right);
 * obj.RemoveRange(left,right);
 */