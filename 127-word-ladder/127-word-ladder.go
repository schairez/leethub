
func ladderLength(beginWord string, endWord string, wordList []string) int {   
    wordSet := make(map[string]struct{}, len(wordList))
    for i := range wordList {
        wordSet[wordList[i]] = struct{}{}
    }
    lvl := 0
    var node string
    queue := make([]string, 0, len(wordList) + 1)
    queue = append(queue, beginWord)
    for len(queue) != 0 {
        lvl++
        for currLen := len(queue); currLen != 0; currLen-- {
            node, queue = queue[0], queue[1:]
            if node == endWord {
                return lvl
            }
            for wordKey := range wordSet {
                if canTransform(node, wordKey) {
                    queue = append(queue, wordKey)
                    delete(wordSet, wordKey)
                } 
            }
        }
    }
    return 0
}

func canTransform(w1, w2 string) bool {
    differs := 0
    for i := 0; i < len(w1); i++ {
        if w1[i] != w2[i] {
            differs++
            if differs > 1 {
                return false
            }
        }
    }
    return differs == 1
}