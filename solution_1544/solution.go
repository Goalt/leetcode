package solution_1544

import (
	"fmt"
	"unicode"
)

func makeGood(s string) string {
	if (len(s) == 0) || (len(s) == 1) {
		return s
	}

	res := []byte(s)
	i := 0
	for i < len(res)-1 {
		if unicode.ToLower(rune(res[i])) == unicode.ToLower(rune(res[i+1])) {
			if (unicode.IsUpper(rune(res[i])) && unicode.IsLower(rune(res[i+1]))) ||
				(unicode.IsLower(rune(res[i])) && unicode.IsUpper(rune(res[i+1]))) {
				res = append(res[:i], res[i+2:]...)
				if i != 0 {
					i -= 1
				}
				continue
			}
		}

		i += 1
	}

	return string(res)
}

func Solution() {
	s := "leEeetcode"
	fmt.Println(makeGood(s))
}
