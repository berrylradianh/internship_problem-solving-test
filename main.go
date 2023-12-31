package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"internship-gits/function"
)

func main() {
	fmt.Println("Pilih fungsi:")
	fmt.Println("1. A000124 of Soloanes OEIS")
	fmt.Println("2. Dense Ranking")
	fmt.Println("3. Balanced Bracket")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan pilihan: ")
	menu, _ := reader.ReadString('\n')
	menu = strings.TrimSpace(menu)

	switch menu {
	case "1":
		fmt.Print("Masukkan Jumlah output : ")
		qty, _ := reader.ReadString('\n')
		qty = strings.TrimSpace(qty)
		n, _ := strconv.Atoi(qty)
		fmt.Println(function.Soloanes(n))
	case "2":
		scores := []int{100, 100, 50, 40, 40, 20, 10}
		gitsScores := []int{5, 25, 50, 120}
		result := function.DenseRanking(scores, gitsScores)
		fmt.Println("Git Ranks", result)

		fmt.Println()

		scores = []int{80, 100, 50, 40, 40, 70, 10}
		gitsScores = []int{7, 21, 40, 170}
		result = function.DenseRanking(scores, gitsScores)
		fmt.Println("Git Ranks", result)

		fmt.Println()

		scores = []int{90, 100, 15, 40, 40, 20, 10}
		gitsScores = []int{5, 25, 110}
		result = function.DenseRanking(scores, gitsScores)
		fmt.Println("Git Ranks", result)
	case "3":
		input1 := "{[()]}"
		fmt.Println("Input 1:", input1)
		fmt.Println("Output 1:", function.BalancedBracket(input1))

		fmt.Println()

		input2 := "{[(])}"
		fmt.Println("Input 2:", input2)
		fmt.Println("Output 2:", function.BalancedBracket(input2))

		fmt.Println()

		input3 := "{(([])[])[]}"
		fmt.Println("Input 3:", input3)
		fmt.Println("Output 3:", function.BalancedBracket(input3)) // Output: YES
	default:
		fmt.Println("Input yang anda masukkan salah!")
	}
}
