package solution_1957

func makeFancyString(s string) string {
	type counter struct {
		Symbol byte
		Count  int
	}

	if s == "" {
		return ""
	}

	countsMas := []counter{{Symbol: s[0], Count: 1}}
	i := 1
	j := 1

	for i < len(s) {
		if s[i] == countsMas[j-1].Symbol {
			countsMas[j-1].Count += 1
		} else {
			countsMas = append(countsMas, counter{Count: 1, Symbol: s[i]})
			j += 1
		}

		i += 1
	}

	var result []byte
	for _, v := range countsMas {
		for i := 0; (i < v.Count) && (i < 2); i += 1 {
			result = append(result, v.Symbol)
		}
	}

	return string(result)
}

func Solution() {
	s := "aab"

	print(makeFancyString(s))
}
