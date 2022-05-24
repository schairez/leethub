// 20. Valid Parentheses
// time: O(n)
// space: O(n)

impl Solution {
    pub fn is_valid(s: String) -> bool {
        let mut stack: Vec<char> = vec![];
        for ch in s.chars() {
            match ch {
                '(' => stack.push(')'),
                '[' => stack.push(']'),
                '{' => stack.push('}'),
                '}' | ']' | ')' if stack.pop() != Some(ch) => {
                    return false
                },
                _ => {},
            }
        }
    return stack.is_empty()
    }
}