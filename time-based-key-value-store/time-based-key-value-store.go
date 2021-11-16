
type Pair struct {
    Timestamp int
    Value string
}
type Pairs []Pair

type TimeMap struct {
    KVStore map[string]Pairs
}

func Constructor() TimeMap {
    return TimeMap{KVStore: make(map[string]Pairs)}
}


func (this *TimeMap) Set(key string, value string, timestamp int)  {
    if _, ok := this.KVStore[key]; !ok {
        this.KVStore[key] = make(Pairs, 0)
    } 
    this.KVStore[key] = append(this.KVStore[key], Pair{Timestamp:timestamp, Value:value} )
}

/*
All the timestamps timestamp of set are strictly increasing.
therefore a candidate for binSearch
*/

func (this *TimeMap) Get(key string, timestamp int) string {
    binSearch := func(v Pairs, target int) string {
        lo, hi := 0, len(v)
        for lo < hi {
            mid := lo + (hi-lo) >> 1
            if v[mid].Timestamp <= target {
                lo = mid + 1
            } else {
                hi = mid
            }
        }
        if hi == 0 { 
            return ""
        } 
        return v[hi-1].Value
    }
    vals, ok := this.KVStore[key] 
    if !ok { return "" }
    
    return binSearch(vals, timestamp)
}


/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */