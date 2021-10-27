package solution_171

import "fmt"

func titleToNumber(columnTitle string) int {
	var res int

	por := 1
	for i := len(columnTitle) - 1; i >= 0; i -= 1 {
		res += int(columnTitle[i]-'A'+1) * por
		por *= 26
	}

	return res
}

func Solution() {
	s := "FXSHRXW"
	fmt.Println(titleToNumber(s))
}
