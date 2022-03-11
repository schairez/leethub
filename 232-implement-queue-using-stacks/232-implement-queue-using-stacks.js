
var MyQueue = function() {
    this.inStack = []
    this.outStack = []
};

/** 
 * @param {number} x
 * @return {void}
 */
MyQueue.prototype.push = function(x) {
    this.inStack.push(x)
};

/**
 * @return {number}
 */
MyQueue.prototype.pop = function() {
    this.peek()
    n = this.outStack.length
    v = this.outStack[n-1]
    this.outStack.pop()
    return v
};

/**
 * @return {number}
 */
MyQueue.prototype.peek = function() {
    if (this.outStack.length == 0) {
        while (this.inStack.length != 0) {
            v = this.inStack.pop()
            this.outStack.push(v)
        }
    }
    n = this.outStack.length
    return this.outStack[n-1]
};

/**
 * @return {boolean}
 */
MyQueue.prototype.empty = function() {
    return this.inStack.length == 0 && this.outStack.length == 0
};




/** 
 * Your MyQueue object will be instantiated and called as such:
 * var obj = new MyQueue()
 * obj.push(x)
 * var param_2 = obj.pop()
 * var param_3 = obj.peek()
 * var param_4 = obj.empty()
 */