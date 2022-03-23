
// 818. Race Car

// naive BFS w/o pruning would lead to O(2^n) runtime and space;
// but here we're not enumerating the accelerate and reverse option for
// each node; with the pruning step our bound reduces dramatically
// as it takes some logN steps to overshoot our target and some additional
// n steps hard cap worst case to find the shortest distance from here
// time: O(nlogn)
// space: O(nlogn)

func racecar(target int) int {
    type pair struct {pos, vel int}
    queue := make([]pair, 0, target >> 1)
    var (
        currNode pair
        srcNode pair
    )
    srcNode = pair{0,1}
    queue = append(queue, srcNode)
    lvl := 0 
    for len(queue) != 0 {
        for n := len(queue); n != 0; n-- {
            currNode, queue = queue[0], queue[1:]   
            if currNode.pos == target {
                return lvl
            }
            pos, vel := currNode.pos, currNode.vel
            if pos + vel <= 2 * target && pos + vel > 0 {
                queue = append(queue, pair{pos+vel, vel*2})
            }
            // if we overshoot
            if pos + vel > target && vel > 0 {
                queue = append(queue, pair{pos, -1})
            } else if pos + vel < target && vel < 0 { // if we get further from t and neg vel
                queue = append(queue, pair{pos, 1})
            }
        }
        lvl++
    }
    
    return -1
}