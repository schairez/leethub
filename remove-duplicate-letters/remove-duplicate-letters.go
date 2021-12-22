import "strings"

//time: O(n)
//space: O(1)

func removeDuplicateLetters(s string) string {
    getIdx := func(b byte) int {
        return int(b - 'a')
    }
    var st stack
    var (
        visited [26]bool
        charsFreq [26]int
    )
    for i := range s {
        charsFreq[getIdx(s[i])]++
    }
    for i := range s {
        currCharIdx := getIdx(s[i])
        charsFreq[currCharIdx]--
        if visited[currCharIdx] == true {
            continue
        }
        for !st.isEmpty() &&
        getIdx(s[st.peek()]) > currCharIdx &&
        charsFreq[getIdx(s[st.peek()])] != 0 {
            visited[getIdx(s[st.pop()])] = false
        }
        visited[currCharIdx] = true
        st.push(i)
    }
    var sb strings.Builder
    sb.Grow(len(st))
    for i := 0; i < len(st); i++ {
        sb.WriteByte(s[st[i]])
    }
    return sb.String()
}


type stack []int

func (st *stack) isEmpty() bool {
    return len((*st)) == 0
}

func (st *stack) push(val int) {
    (*st) = append((*st), val)
}

func (st *stack) peek() int {
    return (*st)[len((*st))-1]
}
func (st *stack) pop() int {
    val := st.peek()
    (*st) = (*st)[:len((*st))-1]
    return val
}
