type Logger struct {
    KVStore map[string]int
}


func Constructor() Logger {
    return Logger{ make(map[string]int) }
}


func (this *Logger) ShouldPrintMessage(timestamp int, message string) bool {
    if allowedTime, ok := this.KVStore[message]; ok {
        if timestamp < allowedTime {
            return false
        } 
        this.KVStore[message] = timestamp + 10
        return true
    }
    this.KVStore[message] = timestamp + 10
    return true
}


/**
 * Your Logger object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.ShouldPrintMessage(timestamp,message);
 */