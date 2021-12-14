//time: O(n)
//space: O(n)

func removeKdigits(num string, k int) string {
    if k == len(num) {
        return "0"
    }
    var deque []byte
    for i:=0; i < len(num); i++ {
        for k != 0 && len(deque) > 0 && deque[len(deque)-1] > num[i] {
            deque = deque[:len(deque)-1]
            k--
        }
        deque = append(deque, num[i])
    }  
    deque = deque[:len(deque)-k]
    
    for len(deque) > 1 && deque[0] == '0' {
        deque = deque[1:]
    }
    
    return string(deque)
}