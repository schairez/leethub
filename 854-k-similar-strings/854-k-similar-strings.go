
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
    /*
    for i := 0; i < n; i++ {
        idx := int(s2[i]-'a')
        if charLoc[idx] == nil {
            charLoc[idx] = make(map[int]struct{})
        }
        charLoc[idx][i] = struct{}{} 
    }
    */
    
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
            nIdx := i+1
            for nIdx < n {
                if node[nIdx] == s2[nIdx] || node[nIdx] != s2[i] {
                    nIdx++
                    continue
                }
                node[i], node[nIdx] = node[nIdx], node[i]
                permKey := string(node) 
                if _, exists := visited[permKey]; !exists {
                    queue = append(queue, append([]byte{}, node...))
                    visited[permKey] = struct{}{}
                }
                node[i], node[nIdx] = node[nIdx], node[i]
                nIdx++
            }
            /*
                if node[i] == s2[i] {
                    continue
                }
                chIdx := int(s1[i]-'a')
                // find correct loc for char
                for corrIdx := range charLoc[chIdx] {
                    if node[corrIdx] == node[i] {
                        continue
                    }
                    node[i], node[corrIdx] = node[corrIdx], node[i]
                    
                    permKey := string(node) 
                    if _, exists := visited[permKey]; ! exists {
                        queue = append(queue, append([]byte{}, node...))
                        visited[permKey] = struct{}{}
                    }
                    node[i], node[corrIdx] = node[corrIdx], node[i]
                }
            }
            */
        }
        lvl++
    }
    return -1
    
}

