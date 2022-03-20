func min(a, b int) int {if a <=b { return a}; return b}

func racecar(target int) int {
    //maxInt32 := (1 << 31) - 1
    type nodeData struct { position, velocity, steps int }
    var queue []nodeData
    var (
        currNode nodeData
        sourceNode nodeData
    )
    sourceNode = nodeData{0, 1, 0}
    queue = append(queue, sourceNode)
    //res := maxInt32
    
    // 
    /////  {1, 2, 1}
    //  {3, 4, 2}
    //    {3, -1, 3}
    ///////// {3, 1, 4} // AARRA
    ////////  {4, 2, 5}
    for len(queue) != 0 {
        currNode, queue = queue[0], queue[1:]
        if currNode.position == target {
            //res = min(res, currNode.steps)
            return currNode.steps
        }
        pos, vel, steps := currNode.position, currNode.velocity, currNode.steps
        queue = append(queue, nodeData{pos+vel, vel*2, steps+1})
        
        if (pos + vel > target && vel > 0 ) || (pos > target && vel > 0) {
            queue = append(queue, nodeData{pos, -1, steps+1})
        } else if (pos + vel < target  && vel < 0) || (pos < target && vel < 0) {
            queue = append(queue, nodeData{pos, 1, steps+1})
        }
        
    }
    
    return -1
    //return res
}