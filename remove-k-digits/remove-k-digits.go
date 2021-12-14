//time: O(n)
//space: O(n)

func removeKdigits(num string, k int) string {
    var deque []byte
    for i:=0; i < len(num); i++ {
        for k > 0 && len(deque) > 0 && deque[len(deque)-1] > num[i] {
            deque = deque[:len(deque)-1]
            k--
        }
        deque = append(deque, num[i])
    }  
    for k != 0 {
        deque = deque[:len(deque)-1]
        k--
    }
    for len(deque) !=0 && deque[0] == '0' {
        deque = deque[1:]
    }
    if len(deque) == 0 {
        return "0"
    }
    
    return string(deque)
}