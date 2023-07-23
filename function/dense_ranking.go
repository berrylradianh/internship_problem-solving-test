package function

import (
	"fmt"
	"sort"
)

// Buat fungsi yang digunakan untuk menyelesaikan permasalahan Dense Ranking!

func DenseRanking(scores []int, gitsScores []int) []int {
	sort.Slice(scores, func(i, j int) bool {
		return scores[j] < scores[i]
	})

	var rankScores []int
	value := 1
	for i, score := range scores {
		if i == 0 || score == scores[i-1] {
			rankScores = append(rankScores, value)
		} else {
			value++
			rankScores = append(rankScores, value)
		}
	}

	rankMap := make(map[int]int)
	rank := 1
	for i, score := range scores {
		if i > 0 && score != scores[i-1] {
			rank++
		}
		rankMap[score] = rank
	}

	var gitsRanks []int
	for i := 0; i < len(gitsScores); i++ {
		for j := 0; j < len(scores); j++ {
			if gitsScores[i] == scores[j] {
				gitsRanks = append(gitsRanks, rankMap[gitsScores[i]])
				break
			} else if gitsScores[i] > scores[j] {
				gitsRanks = append(gitsRanks, rankScores[j])
				break
			} else if gitsScores[i] < scores[len(scores)-1] {
				gitsRanks = append(gitsRanks, rankScores[len(rankScores)-1]+1)
				break
			}
		}
	}

	fmt.Println("Scores & Ranks {", rankMap, "}")
	fmt.Println("Git Scores", gitsScores)

	return gitsRanks
}
