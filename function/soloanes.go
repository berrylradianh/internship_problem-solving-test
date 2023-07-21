package function

import "fmt"

func Soloanes(n int) {
	var result []int

	for i := 0; i < n; i++ {
		formula := i*(i+1)/2 + 1
		result = append(result, formula)
	}

	fmt.Println(result)
}
