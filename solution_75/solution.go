package solution_75

import "fmt"

func sortColors(nums []int) {
	var cntRed, cntWhite, cntBlue int
	for _, v := range nums {
		switch v {
		case 0:
			cntRed += 1
		case 1:
			cntWhite += 1
		case 2:
			cntBlue += 1
		}
	}

	var i int

	for j := 0; j < cntRed; j += 1 {
		nums[i] = 0
		i += 1
	}

	for j := 0; j < cntWhite; j += 1 {
		nums[i] = 1
		i += 1
	}

	for j := 0; j < cntBlue; j += 1 {
		nums[i] = 2
		i += 1
	}
}

func Solution() {
	nums := []int{1}
	sortColors(nums)
	fmt.Println(nums)
}
