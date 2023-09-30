/**
 * The read4 API is already defined for you.
 *
 *     read4 := func(buf4 []byte) int
 *
 * // Below is an example of how the read4 API can be called.
 * file := File("abcdefghijk") // File is "abcdefghijk", initially file pointer (fp) points to 'a'
 * buf4 := make([]byte, 4) // Create buffer with enough space to store characters
 * read4(buf4) // read4 returns 4. Now buf = ['a','b','c','d'], fp points to 'e'
 * read4(buf4) // read4 returns 4. Now buf = ['e','f','g','h'], fp points to 'i'
 * read4(buf4) // read4 returns 3. Now buf = ['i','j','k',...], fp points to end of file
 */

var solution = func(read4 func([]byte) int) func([]byte, int) int {
    // implement read below.
    buf4 := make([]byte, 4)
    read, position := 0, 0
    return func(buf []byte, n int) int {
        numCharsRead := 0
        for numCharsRead < n {
            if read == 0 {
                read = read4(buf4)
                position = 0
                fmt.Println(read)
            }
            if read == 0 {
                break
            }
            buf[numCharsRead] = buf4[position]
            numCharsRead++
            position++
            read--
        }
        return numCharsRead
    }
}