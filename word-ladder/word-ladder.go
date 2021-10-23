// time: O(N^2* len(word));
// space: O(n)

type Node struct {
  word string
  dist int
}

type set map[string]bool

func NewSet(words []string) set {
  s := make(set)
  for i := range words {
      s[words[i]] = true
  }
    return s
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
    n := len(beginWord)
    //curr := beginWord
    s := NewSet(wordList)
    //precondition
    if _, ok := s[endWord]; !ok {
        return 0
    }
    q := []*Node{}
    q = append(q, &Node{beginWord, 1})
    var degrees int
    var dist int
    for len(q) > 0 {
        node := q[0] //hit
        if node.word == endWord { return node.dist }
        dist = node.dist
        curr := node.word
        q = q[1:]
        for k := range s { //O(n)
            degrees = 0
            for j := range curr {
                if k[j] == curr[j] { //O(len(word))
                    degrees++
                }
            }
            if degrees == n-1 {
                q = append(q, &Node{k, dist + 1})
                delete(s, k) //O(1)
            }
        }
    }
    return 0
}
