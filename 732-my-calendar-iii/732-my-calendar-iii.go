type val struct { a, b int }
type tree map[int]val


type MyCalendarThree struct {
    t tree
}


func Constructor() MyCalendarThree {
    return MyCalendarThree{t: map[int]val{}}
}


func (this *MyCalendarThree) Book(start int, end int) int {
    var update func(i, s, t int)
    update = func(i, s, t int) {
        if end < s || t < start {
            return
        }
        if start <= s && t <= end {
            v := this.t[i]
            v.a++
            v.b++
            this.t[i] = v
            return
        }
        m := (s+t) / 2
        update(i*2, s, m)
        update(i*2+1, m+1, t)
        v := this.t[i]
        v.a = v.b + max(this.t[i*2].a, this.t[i*2+1].a)
        this.t[i] = v
    }
    end--
    update(1, 0, 1e9)
    return this.t[1].a
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

/**
 * Your MyCalendarThree object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */