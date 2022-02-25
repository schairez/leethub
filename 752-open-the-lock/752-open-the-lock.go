
func openLock(deadends []string, target string) int {
    slots := newWheelSlots()
    var dstNode [4]byte
    var srcNode [4]byte
    for i := 0; i < 4; i++ {
        dstNode[i] = target[i]
        srcNode[i] = '0'
    }
    deadendMap := make(map[[4]byte]struct{}, len(deadends))
    for _, deadendNode := range deadends {
        var key [4]byte
        for j := 0; j < 4; j++ {
            key[j] = deadendNode[j]
        }
        deadendMap[key] = struct{}{}
    }
    // edge case: we need to check if our srcNode is in our deadendsMap
    if _, ok := deadendMap[srcNode]; ok {
        return -1
    }
    
    equalsDstNodeFn := equalsTargetNode(dstNode)
    var (
        queue    [][4]byte
        pollNode [4]byte
        lvl      int  
        nextPrev [2]byte  //turning wheel slot forward and backward by 1 val
    )
    queue = append(queue, srcNode)
    visited := make(map[[4]byte]struct{})
    visited[srcNode] = struct{}{}
    for len(queue) != 0 {
        for currLen := len(queue); currLen != 0; currLen-- {
            pollNode, queue = queue[0], queue[1:]
            if equalsDstNodeFn(pollNode) {
                return lvl
            } 
            for i := 0; i < 4; i++ {
                nextPrev  = slots.nextAndPrev(pollNode[i])
                for j := 0; j < 2; j++ {
                    neiNode := pollNode
                    neiNode[i] = nextPrev[j]
                    if _, ok := deadendMap[neiNode]; ok {
                        continue
                    }
                    if _, ok := visited[neiNode]; ok {
                        continue
                    }
                    queue = append(queue, neiNode)
                    visited[neiNode] = struct{}{}
                }
            }
        }
        lvl++
    }
    return -1
}

func equalsTargetNode(target [4]byte) func(combo [4]byte) bool {
    return func(combo [4]byte) bool {
        for i := 0; i < 4; i++ {
            if target[i] != combo[i] {
                return false
            }
        }
        return true
    }
}


type wheelSlots struct {
    nodes [10]byte
}
func newWheelSlots() wheelSlots {
    nodes := [10]byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
    return wheelSlots{nodes}
}
func (sl wheelSlots) nextAndPrev(currByte byte) [2]byte {
    idx := currByte - '0'
    var next, prev byte
    next = sl.nodes[(idx+1) % 10]
    prev = sl.nodes[(idx-1) % 10]
    if idx == 0 {
        prev = sl.nodes[9]
    }
    return [2]byte{next, prev}
}

/*

func (sl wheelSlots) isValid(idx int) bool {
    return idx >= 0 && idx < 10
}
*/

