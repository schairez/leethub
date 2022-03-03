
type Node struct {
    key int
    val int
    prev, next *Node
}

type LRUCache struct {
    capacity int
    head, tail *Node
    nodeKeyMap map[int]*Node
}


func Constructor(capacity int) LRUCache {
    head, tail := &Node{0,0, nil, nil}, &Node{0,0, nil, nil}
    head.next = tail
    tail.prev = head
    return LRUCache{capacity: capacity,
                    head: head,
                    tail: tail,
                    nodeKeyMap: make(map[int]*Node),
                   }
}


func (this *LRUCache) Get(key int) int {
    res := -1
    node, exists := this.nodeKeyMap[key] 
    if exists {
        res = node.val
        this.deleteNode(node)
        this.addToHead(node)
    }
    return res
}

func (this *LRUCache) deleteNode(node *Node) {
    
    oldPrev := node.prev
    oldNext := node.next
    oldPrev.next = oldNext
    oldNext.prev = oldPrev
    
    /*
    oldNext := node.next
    oldPrev := node.prev
    oldPrev.next = oldNext
    oldNext.prev = oldPrev
    */
}


func (this *LRUCache) addToHead(node *Node) {
    //
    node.prev = this.head
    //
    node.next = this.head.next
    this.head.next = node
    node.next.prev = node
    
} 
//upsert
func (this *LRUCache) Put(key int, value int)  {
    node, exists :=this.nodeKeyMap[key]
    if exists {
        node.val = value
        this.deleteNode(node)
        this.addToHead(node)
        return
    } else {
        newNode := &Node{key: key, val: value}
        this.nodeKeyMap[key] = newNode
        this.addToHead(newNode)
    }
    if len(this.nodeKeyMap) > this.capacity {
        delNode := this.tail.prev
        delete(this.nodeKeyMap, delNode.key)
        this.deleteNode(delNode)
    }
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */