
type Node struct {
    cnt int
    keySet map[string]struct{}
    prev, next *Node
}
//condition: no nil inputs
func (currNode *Node) InsertNode(newNode *Node) {
    if currNode == nil { return }
    nodeAfterNewNode := currNode.next
    newNode.prev = currNode
    currNode.next = newNode
    newNode.next = nodeAfterNewNode
    nodeAfterNewNode.prev = newNode
}

func DeleteNode(delNode *Node) {
    prev := delNode.prev
    next := delNode.next
    prev.next = next
    next.prev = prev
}


type AllOne struct {
    nodeKeySet map[string]*Node
    nodeListHeader *Node
}



func Constructor() AllOne {
    list := &Node{}
    list.cnt = 0
    list.keySet = make(map[string]struct{})
    list.prev = list
    list.next = list
    return AllOne{nodeKeySet: make(map[string]*Node),
                  nodeListHeader: list,
                 }
}


func (this *AllOne) Inc(key string)  {
    if oldNode, exists := this.nodeKeySet[key]; exists {
        if oldNode.next.cnt != oldNode.cnt + 1 {
            oldNode.InsertNode(&Node{
                cnt: oldNode.cnt + 1,
                keySet: make(map[string]struct{}),
            })
        }
        newNode := oldNode.next
        newNode.keySet[key] = struct{}{}
        this.nodeKeySet[key] = newNode
        delete(oldNode.keySet, key)
        if len(oldNode.keySet) == 0 {
            DeleteNode(oldNode)
        }
    } else {
        if this.nodeListHeader.next.cnt != 1 {
            this.nodeListHeader.InsertNode(&Node{
                cnt: 1,
                keySet: make(map[string]struct{}),
            })
        }
        newNode := this.nodeListHeader.next 
        newNode.keySet[key] = struct{}{}
        this.nodeKeySet[key] = newNode
    }
}


func (this *AllOne) Dec(key string)  {
    if oldNode, exists := this.nodeKeySet[key]; exists {
        if oldNode.cnt == 1 {
            delete(this.nodeKeySet, key)
        } else {
            if oldNode.prev.cnt != oldNode.cnt - 1 {
                oldNode.prev.InsertNode(&Node{
                    cnt: oldNode.cnt -1,
                    keySet: make(map[string]struct{}),
                })
            }
            newNode := oldNode.prev
            newNode.keySet[key] = struct{}{}
            this.nodeKeySet[key] = newNode
        }
        delete(oldNode.keySet, key)
        if len(oldNode.keySet) == 0 {
            DeleteNode(oldNode)
        }
    }
}

func (this *AllOne) GetMaxKey() string {
    for key, _ := range this.nodeListHeader.prev.keySet {
        return key
    }
    return ""
}


func (this *AllOne) GetMinKey() string {
    for key, _ := range this.nodeListHeader.next.keySet {
        return key
    }
    return ""
}


/**
 * Your AllOne object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Inc(key);
 * obj.Dec(key);
 * param_3 := obj.GetMaxKey();
 * param_4 := obj.GetMinKey();
 */