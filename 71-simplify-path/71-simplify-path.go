import "strings"

//time: O(n)
//space: O(n)

func simplifyPath(path string) string {
    var (
        dirPath []string
        deque []string
        sb strings.Builder
    )
    dirPath = strings.Split(path, "/")
    sb.Grow(len(dirPath))
    for _, dir := range dirPath {
        if dir == "." || dir == "" {
            continue
        } else if dir == ".." {
            if len(deque) != 0 {
                deque = deque[:len(deque)-1]
            }
        } else {
            deque = append(deque, dir)
        }
    }
    var poppedItem string
    for len(deque) != 0 {
        sb.WriteByte('/')
        poppedItem, deque = deque[0], deque[1:]
        sb.WriteString(poppedItem)
    }
    res := sb.String()
    if len(res) == 0 {
        return "/"
    }
    return res
}
