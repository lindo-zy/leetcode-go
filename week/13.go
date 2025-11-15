package main

func main() {

}

func romanToInt(s string) int {
	numMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	num := 0
	for i := 0; i < len(s); i++ {
		if i < len(s)-1 && numMap[s[i]] < numMap[s[i+1]] {
			num -= numMap[s[i]]
		} else {
			num += numMap[s[i]]
		}
	}
	if 1 <= num && num <= 3999 {
		return num
	}
	return 0
}
