/*
space: O(maxSize)
pop O(1) space & time
incr time: O(k) k elem operations 
*/

type CustomStack struct {
    stack []int
    maxSize int
    //curLen int
    
    
}


func Constructor(maxSize int) CustomStack {
    stack := make([]int, 0, maxSize)
    return CustomStack{stack, maxSize}
    
}


func (this *CustomStack) Push(x int)  {
    if len(this.stack) < this.maxSize {
        this.stack = append(this.stack, x)
    }
    return 
}


func (this *CustomStack) Pop() int {
    if len(this.stack) != 0 {
        val := this.stack[len(this.stack)-1]
        this.stack = this.stack[:len(this.stack)-1]
        return val
        
    }
    return -1 
}


func (this *CustomStack) Increment(k int, val int)  {
        last := min(k, len(this.stack)) 
    
        for i := 0; i < last; i ++ {
            this.stack[i] = this.stack[i] + val
        }
    
}

func min(a, b int) int {
    if a <= b {
        return a 
    }
    return b 
}


/**
 * Your CustomStack object will be instantiated and called as such:
 * obj := Constructor(maxSize);
 * obj.Push(x);
 * param_2 := obj.Pop();
 * obj.Increment(k,val);
 */