package main

func getSmallestString(s string) string {
	n := len(s)
	bs := []byte(s)
	for i := 1; i < n; i++ {
		a, b := bs[i-1], bs[i]
		if a%2 == b%2 && a > b {
			bs[i-1], bs[i] = bs[i], bs[i-1]
			break
		}
	}

	return string(bs)
}
func main() {
	getSmallestString("43520")
}
