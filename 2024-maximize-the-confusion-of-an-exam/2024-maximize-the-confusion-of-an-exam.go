
// 2024. Maximize the Confusion of an Exam
// time: O(2n) ≈ O(n)
// space: O(1)

func max(a, b int) int { if a >= b { return a}; return b}

func maxConsecutiveAnswers(answerKey string, k int) int {
    return max( maxByType(answerKey, k, true), maxByType(answerKey, k, false))
}

func maxByType(answerKey string, k int, isTrueVersion bool) int {
    n := len(answerKey)
    // if isTrueVersion we want to flip the 'F' else flip 'T'
    var answerToFlip byte = 'F'
    if !isTrueVersion {
        answerToFlip = 'T'
    }
    start, end := 0, 0
    res := 0
    numToFlip := 0
    for end < n {
        if answerKey[end] == answerToFlip {
            numToFlip++
        }
        for numToFlip > k {
            if answerKey[start] == answerToFlip {
                numToFlip--
            }
            start++
        }
        res = max(res, end - start +1)
        end++
    }
    return res
} 


