package day07

import "sort"

func SolvePartB(input []string) int {
	handsWithBids := make([]handWithBid, len(input))

	for i, line := range input {
		handsWithBids[i] = parseHandWithBid(line, true)
	}

	sort.Slice(handsWithBids, func(i, j int) bool {
		return compareHands(handsWithBids[i].hand, handsWithBids[j].hand, true) == -1
	})

	winnings := 0

	for i := 0; i < len(handsWithBids); i++ {
		winnings += handsWithBids[i].bid * (i + 1)
	}

	return winnings
}
