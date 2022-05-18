
func min(a, b int) int {if a <= b { return a}; return b}


func alienOrder(words []string) string {
    n := len(words)
    // constraint: words[i] consists of only lowercase English letters.
    var (
        graph    [26][]byte
        inDegree [26]int
        exists   [26]bool
        queue    []byte
        node     byte
    )
    numNodes := 0
    for _, word := range words {
        for i := 0; i < len(word); i++ {
            if !exists[word[i]-'a'] {
                exists[word[i]-'a'] = true
                numNodes++
            }
        }
    }
    for i := 1; i < n; i++ { // O(numWords)
        w1, w2 := words[i-1], words[i]
        n1, n2 := len(w1), len(w2)
        if n2 < n1 && string(w1[:n2]) == w2 { // O(lenWord2)
            return ""
        }
        minLen := min(n1, n2)
        for j := 0; j < minLen; j++ { //O(minLen)
            if w1[j] != w2[j] {
                inDegree[w2[j]-'a']++
                graph[w1[j]-'a'] = append(graph[w1[j]-'a'], w2[j])
                break
            }
        }
    }
    for i := 0; i < 26; i++ {
        if !exists[i] {
            continue
        }
        if inDegree[i] == 0 {
            node = byte(i + 'a')
            queue = append(queue, node)
        }
    }
    var sb strings.Builder
    for len(queue) != 0 {
        for currLen := len(queue); currLen != 0; currLen-- {
            node, queue = queue[0], queue[1:]
            sb.WriteByte(node)
            neighbors := graph[node - 'a']
            if len(neighbors) != 0 {
                for _, nei := range neighbors {
                    inDegree[nei - 'a']--
                    if inDegree[nei - 'a'] == 0 {
                        queue = append(queue, nei)
                    }
                } 
            }
        }
    }
    res := sb.String()
    if len(res) != numNodes {
        return ""
    } 
    return sb.String()
}






