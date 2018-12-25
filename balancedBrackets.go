package main

var opening = map[rune]bool {
	'[': true,
	'(': true,
	'{': true,
}

var closing = map[rune]bool {
	']': true,
	')': true,
	'}': true,
}

var maching = map[rune]rune {
	'[':']',
	'(':')',
	'{':'}',
}

func BalancedBrackets(s string) bool {
	stack := []rune{}
	for _, v := range s {
		if _, ok := opening[v]; ok {
			stack = append(stack, v)
			continue
		}
		if _, ok := closing[v]; ok {
			if len(stack) == 0 {
				return false
			}
			if stack[len(stack)-1] == maching[v] {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}

func main() {

}