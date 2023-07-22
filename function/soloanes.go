package function

// Buat fungsi untuk menyelesaikan rumus A000124 of Sloane’s OEIS!

func Soloanes(n int) []int {
	var result []int

	for i := 0; i < n; i++ {
		formula := i*(i+1)/2 + 1
		result = append(result, formula)
	}

	return result
}
