func racecar(target int) int {
    type pair struct {pos, vel int}
    var queue []pair
    var (
        currNode pair
        srcNode pair
    )
    srcNode = pair{0,1}
    queue = append(queue, srcNode)
    
    dist := 0 
    for len(queue) != 0 {
        //n := len(queue)
        for n := len(queue); n != 0; n-- {
            currNode, queue = queue[0], queue[1:]   
            if currNode.pos == target {
                return dist
            }
            pos, vel := currNode.pos, currNode.vel
            queue = append(queue, pair{pos+vel, vel*2})
            if pos + vel > target && vel > 0 {
                queue = append(queue, pair{pos, -1})
            } else if pos + vel < target && vel < 0 {
                queue = append(queue, pair{pos, 1})
            }
        }
        dist++
    }
    
    return -1
}