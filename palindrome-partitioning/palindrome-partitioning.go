//time: O(n*2^n)
//space: O(n)

func partition(s string) [][]string {
	res := [][]string{}
	cur := make([]string, 0, len(s))
	dfs(s, 0, cur, &res)
	return res
}

//O(n*2^n)
func dfs(s string, i int, cur []string, res *[][]string) {
	if len(s) == i {
		tmp := make([]string, len(cur))
		copy(tmp, cur)
		*res = append(*res, tmp)
	}
	for j := i; j < len(s); j++ {
		if isPalindrome(s[i : j+1]) {
			dfs(s, j+1, append(cur, s[i:j+1]), res)
		}
	}
}
//O(n)
func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}
