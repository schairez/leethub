// 1087. Brace Expansion
// BFS Approach
// time: O(n) to iter string and append to lvlNodes
// space: O(numLvl*numNodesAtLvl) for lvlNodes store
// for queue BFS
// time: O(|V| + |E|); our n-ary tree varies per input
// we traverse through lvlNodes and append prev lvlNodes to currLvlNode edges
// our branch factor won't be constant but we can 
// sort takes O(nlogn) time and O(logn) space 
// time ≈ O(n*n)
// space ≈ O(n*n)


func isAlpha(b byte) bool {
    return b >= 'a' && b <= 'z'
}
func expand(s string) []string {
    n := len(s)
    if n == 1 {
        return []string{string(s[0])}
    }
    var lvlNodes [][]byte
    idx := 0
    for idx < n {
        if isAlpha(s[idx]) {
            var lvl []byte
            lvl = append(lvl, s[idx])
            lvlNodes = append(lvlNodes, lvl)
            idx++
        } else if s[idx] == '{' {
            var lvl []byte
            for idx < n && s[idx] != '}' {
                if isAlpha(s[idx]) {
                    lvl = append(lvl, s[idx])
                }
                idx++
            }
            lvlNodes = append(lvlNodes, lvl)
        } else {
            idx++
        }
    }
    var (
        res []string
        node []byte
        queue [][]byte
    )
    numLvls := len(lvlNodes)
    lvl := 0
    queue = append(queue, []byte{})
    for len(queue) != 0 {
        for currLen := len(queue); currLen != 0; currLen-- {
            node, queue = queue[0], queue[1:]
            if lvl == numLvls {
                res = append(res, string(node))
            } else {
                numNodesAtLvl := len(lvlNodes[lvl])
                for i := 0; i < numNodesAtLvl; i++ { 
                    node = append(node, lvlNodes[lvl][i])
                    queue = append(queue, append([]byte{}, node...))
                    node = node[:len(node)-1]
                }
            }
        }
        lvl++
    }
    sort.Strings(res)
    return res
}