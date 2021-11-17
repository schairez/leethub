
type Logger struct {
    KVStore map[string]int
}


func Constructor() Logger {
    l := Logger{KVStore: make(map[string]int)}
    return l
}


func (this *Logger) ShouldPrintMessage(timestamp int, message string) bool {
    fmt.Println(this.KVStore)
    switch _, exists := this.KVStore[message]; exists {
    case true:
        if timestamp - this.KVStore[message] >= 10 {
            this.KVStore[message] = timestamp
            return true
        }
        return false
    default: //doesnt exist
        this.KVStore[message] = timestamp
        return true
    }
}


/**
 * Your Logger object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.ShouldPrintMessage(timestamp,message);
 */