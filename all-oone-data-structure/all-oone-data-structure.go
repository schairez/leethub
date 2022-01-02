//time: O(1)
//space: O(n)

type KeyNode struct {
    cnt int
    vals map[string]struct{}
    prev, next *KeyNode
}

type AllOne struct {
    keySet map[string]*KeyNode
    head, tail *KeyNode
}

var maxInt32 int = (1 << 31) - 1

func Constructor() AllOne {
    keySet := make(map[string]*KeyNode)
    head := &KeyNode{cnt: 0}
    tail := &KeyNode{cnt: maxInt32}
    head.next = tail
    tail.prev = head
    return AllOne{keySet: keySet, head: head, tail: tail}
}


func (this *AllOne) Inc(key string)  {
    var keyFreqNode *KeyNode
    if _, exists := this.keySet[key]; !exists {
        keyFreqNode = this.addToHead(key)
    } else {
        keyFreqNode = this.addToNextFreq(key)
    } 
    this.keySet[key] = keyFreqNode
}
func (this *AllOne) removeNode(node *KeyNode) {
    currPrev := node.prev
    currNext := node.next
    node.prev.next = nil
    node.next.prev = nil
    currPrev.next = currNext
    currNext.prev = currPrev
}
//assumes key we're incrementing not in keyset
func (this *AllOne) addToHead(key string) *KeyNode {
    currNextHeadNode := this.head.next 
    if currNextHeadNode.cnt != 1 {
        //currNextNode := this.head.next
        newNextNode := &KeyNode{cnt: 1, vals: make(map[string]struct{})}
        newNextNode.vals[key] = struct{}{}
        newNextNode.next = this.head.next
        this.head.next = newNextNode
        newNextNode.next.prev = newNextNode
        newNextNode.prev = this.head
        return newNextNode
    }
    currNextHeadNode.vals[key] = struct{}{}
    return currNextHeadNode
}
    
func (this *AllOne) addToNextFreq(key string) *KeyNode {
    var returnNode *KeyNode
    currNode := this.keySet[key]
    currNextFreqNode := currNode.next
    currKeyFreq := currNode.cnt
    expNextFreq := currKeyFreq+1
    if currNode.next.cnt != expNextFreq {
        newNextNode := &KeyNode{cnt: expNextFreq, vals: make(map[string]struct{})}
        newNextNode.next = currNode.next
        currNode.next = newNextNode
        newNextNode.next.prev = newNextNode
        newNextNode.prev = currNode
        returnNode = newNextNode
    } else {
        returnNode = currNextFreqNode
    }
    currNode.next.vals[key] = struct{}{}
    delete(currNode.vals, key)
    if len(currNode.vals) == 0 {
        this.removeNode(currNode)
    }
    return returnNode
}
//condition: guaranteed key exists
func (this *AllOne) Dec(key string)  {
    currNode := this.keySet[key]
    prevNode := currNode.prev
    //nextNode := currNode.next
    currNodeFreq := currNode.cnt
    expDecrFreq := currNodeFreq-1
    delete(currNode.vals, key)
    if expDecrFreq == 0 {
        delete(this.keySet, key)
    } else if expDecrFreq == prevNode.cnt {
        prevNode.vals[key] = struct{}{}
        this.keySet[key] = prevNode
    } else {
        newPrevNode := &KeyNode{cnt: expDecrFreq, vals: make(map[string]struct{})}
        newPrevNode.vals[key] = struct{}{}
        this.keySet[key] = newPrevNode
        currNode.prev.next = nil
        currNode.prev = newPrevNode
        newPrevNode.prev = prevNode
        prevNode.next = newPrevNode
        newPrevNode.next = currNode
    }
    if len(currNode.vals) == 0 {
        this.removeNode(currNode)
    }
}


func (this *AllOne) GetMaxKey() string {
    var maxKey string
    sentinel := ""
    maxNode := this.tail.prev
    if maxNode == this.head {
        maxKey = sentinel
    } else {
        for key := range maxNode.vals {
            maxKey = key
            break
        }
    }
    return maxKey
}
func (this *AllOne) GetMinKey() string {
    var minKey string
    sentinel := ""
    minNode := this.head.next
    if minNode == this.tail {
        minKey = sentinel
    } else {
        for key := range minNode.vals {
            minKey = key
            break
        }
    }
    return minKey
}


/**
 * Your AllOne object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Inc(key);
 * obj.Dec(key);
 * param_3 := obj.GetMaxKey();
 * param_4 := obj.GetMinKey();
 */