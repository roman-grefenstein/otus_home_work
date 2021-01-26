package hw03_frequency_analysis //nolint:golint,stylecheck
import (
	"sort"
	"strings"
)

type WordStat struct {
	word  string
	count int
}

func TopN(str string, n int) []string {
	wordStats := make(map[string]int)

	words := strings.Fields(str)

	if len(words) == 0 || n == 0 {
		return []string{}
	}

	for _, word := range words {
		wordStats[word]++
	}

	stats := make([]WordStat, 0, len(words))
	for word, count := range wordStats {
		stats = append(stats, WordStat{word, count})
	}
	sort.SliceStable(stats, func(i, j int) bool {
		return stats[i].count > stats[j].count
	})

	var resultSize int
	if resultSize = n; n > len(stats) {
		resultSize = len(stats)
	}

	result := make([]string, 0, resultSize)
	for _, wordStat := range stats[:resultSize] {
		result = append(result, wordStat.word)
	}
	return result
}

func Top10(str string) []string {
	return TopN(str, 10)
}
