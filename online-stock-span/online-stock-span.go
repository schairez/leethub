//time: O(n)
//space: O(1)

type pair struct {
    price int
    cnt int
}

type StockSpanner struct {
    stack []pair
}


func Constructor() StockSpanner {
    st := []pair{}
    return StockSpanner{stack: st}
}


func (this *StockSpanner) Next(price int) int {
    cnt := 1
    for len(this.stack) > 0 && this.stack[len(this.stack)-1].price <= price {
        prevCnt := this.stack[len(this.stack)-1].cnt
        cnt += prevCnt
        this.stack = this.stack[:len(this.stack)-1]
    }
    this.stack = append(this.stack, pair{price, cnt})
    return cnt
    
}


/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */