package main

func main() {
	ret := isPalindrome(12521)
	println(ret)
}

func isPalindrome(x int) bool {
	if x < 0 || (x % 10 == 0 && x != 0) {
		return false
	}

	var palindrome int = 0
	var j int = x

	for j != 0 {
		temp := palindrome * 10 + j % 10
		palindrome = temp
		j /= 10
	}

	if palindrome == x {
		return true
	}

	return false
}