/*
time: O(n!)
space: O(n)

*/


/*
digits = "23"
Output: ["ad","ae","af","bd","be","bf","cd","ce","cf"]
            

*/

func letterCombinations(digits string) []string {
    var res []string
    if len(digits) == 0 {
        return res
    }
    phoneMap := map[byte]string{
        byte('2'): "abc", byte('3'): "def",
        byte('4'): "ghi", byte('5'): "jkl",
        byte('6'): "mno", byte('7'): "pqrs",
        byte('8'): "tuv", byte('9'): "wxyz",
    }
    combo := make([]byte, len(digits))
    permute(&res, combo, phoneMap, digits, 0)
    return res
}

func permute(res *[]string, combo []byte, phoneMap map[byte]string, digits string, start int) {
    if start == len(digits) {
        tmp := make([]byte, len(digits))
        copy(tmp, combo)
        *res = append(*res, string(tmp))
        return
        
    }
    var letters string
    var key byte
    key = digits[start]
    letters = phoneMap[key]
    for i:=0; i <len(letters); i++ {
        combo[start] = letters[i]
        permute(res, combo, phoneMap, digits, start+1)
        combo[start] = 0
    }
        
}