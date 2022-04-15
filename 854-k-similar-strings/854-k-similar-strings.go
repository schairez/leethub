
// condition:
// s1 and s2 contain only lowercase letters from the set {'a', 'b', 'c', 'd', 'e', 'f'}.
// idx 0-5 with 6 possible chars

func kSimilarity(s1 string, s2 string) int {
    n := len(s1)
    if n == 1 {
        return 0
    }
    if n == 2 { 
        return 1
    }
    var (
        queue   [][]byte
        node    []byte
        //charLoc [6]map[int]struct{}
    )
    
    visited := make(map[string]struct{})
    visited[s1] = struct{}{}
    queue = append(queue, []byte(s1))
    lvl := 0
    for len(queue) != 0 {
        for currLen := len(queue); currLen != 0; currLen-- {
            node, queue = queue[0], queue[1:]
            if string(node) == s2 {
                return lvl
            }
            i := 0
            for i < n && node[i] == s2[i] {
                i++
            }
            j := i+1
            for j < n {
                if node[j] == s2[j] || node[j] != s2[i] {
                    j++
                    continue
                }
                node[i], node[j] = node[j], node[i]
                permKey := string(node) 
                if _, exists := visited[permKey]; !exists {
                    queue = append(queue, append([]byte{}, node...))
                    visited[permKey] = struct{}{}
                }
                node[i], node[j] = node[j], node[i]
                j++
            }
        }
        lvl++
    }
    return -1
    
}

