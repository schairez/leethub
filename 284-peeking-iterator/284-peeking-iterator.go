/*   Below is the interface for Iterator, which is already defined for you.
 *
 *   type Iterator struct {
 *       
 *   }
 *
 *   func (this *Iterator) hasNext() bool {
 *		// Returns true if the iteration has more elements.
 *   }
 *
 *   func (this *Iterator) next() int {
 *		// Returns the next element in the iteration.
 *   }
 */

//we're yielding one elem at a time
//time: O(1)
//space: O(1)

type PeekingIterator struct {
    iter *Iterator
    peekVal int
    hasNextV bool
}

func Constructor(iter *Iterator) *PeekingIterator {
    return &PeekingIterator{
        iter: iter,
        peekVal: iter.next(),
        hasNextV: iter.hasNext(),
    }
}

func (this *PeekingIterator) hasNext() bool {
    return this.hasNextV
}

func (this *PeekingIterator) next() int {
    curr := this.peekVal
    if !this.iter.hasNext() {
        this.hasNextV = false
    } else {
        this.peekVal = this.iter.next() 
    }
    return curr
}

func (this *PeekingIterator) peek() int {
    return this.peekVal
}