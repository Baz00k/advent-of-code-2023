package day07

import (
	"strconv"
)

type CardMap map[rune]int

var cardsMapWithoutJoker = CardMap{
	'2': 1,
	'3': 2,
	'4': 3,
	'5': 4,
	'6': 5,
	'7': 6,
	'8': 7,
	'9': 8,
	'T': 9,
	'J': 10,
	'Q': 11,
	'K': 12,
	'A': 13,
}

var cardsMapWithJoker = CardMap{
	'J': 1,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'T': 10,
	'Q': 11,
	'K': 12,
	'A': 13,
}

const JOKER = 1

type hand struct {
	cards    [5]int
	handType int
}

func handMap(cards [5]int) map[int]int {
	// Create a map between the card and the number of times it appears in the hand
	handMap := make(map[int]int)

	for _, card := range cards {
		handMap[card]++
	}

	return handMap
}

func getHandType(cards [5]int, withJoker bool) int {
	handMap := handMap(cards)
	maxCount := 0
	jokerCount := 0

	if withJoker {
		jokerCount = handMap[JOKER]
		delete(handMap, JOKER)
	}

	for _, count := range handMap {
		if count > maxCount {
			maxCount = count
		}
	}

	// Five of a kind
	if maxCount == 5 || (maxCount+jokerCount) == 5 {
		return 7
	}

	// Four of a kind
	if maxCount == 4 || (maxCount+jokerCount) == 4 {
		return 6
	}

	// Full house
	for k, count := range handMap {
		// We can use joker as a three of a kind
		if (count + jokerCount) == 3 {
			for k2, count := range handMap {
				if k == k2 {
					continue
				}

				if count == 2 {
					return 5
				}
			}
		}

		// We can use joker as a pair
		if (count + jokerCount) == 3 {
			for k2, count := range handMap {
				if k == k2 {
					continue
				}

				if count == 2 {
					return 5
				}
			}
		}
	}

	// Three of a kind
	if maxCount == 3 || (maxCount+jokerCount) == 3 {
		return 4
	}

	// Two pairs
	pairCount := 0
	for _, count := range handMap {
		jokersUsed := false

		if count == 2 {
			pairCount++
			continue
		}

		if !jokersUsed && (count+jokerCount) == 2 {
			pairCount++
			jokersUsed = true
		}
	}
	if pairCount == 2 {
		return 3
	}

	// One pair
	if maxCount == 2 || (maxCount+jokerCount) == 2 {
		return 2
	}

	// High card
	return 1
}

func compareHands(a, b hand, withJoker bool) int {
	if a.handType > b.handType {
		return 1
	}
	if a.handType < b.handType {
		return -1
	}

	// Both hands have the same type
	// The first hand that has a higher card wins
	for i := 0; i < 5; i++ {
		if a.cards[i] > b.cards[i] {
			return 1
		}
		if a.cards[i] < b.cards[i] {
			return -1
		}
	}

	return 0
}

func parseHand(s string, cardMap CardMap, withJoker bool) hand {
	var h hand

	for i := 0; i < 5; i++ {
		h.cards[i] = cardMap[rune(s[i])]
	}

	h.handType = getHandType(h.cards, withJoker)

	return h
}

type handWithBid struct {
	hand hand
	bid  int
}

func parseHandWithBid(s string, withJoker bool) handWithBid {
	var h handWithBid

	if withJoker {
		h.hand = parseHand(s[:5], cardsMapWithJoker, withJoker)
	} else {
		h.hand = parseHand(s[:5], cardsMapWithoutJoker, withJoker)
	}

	bid, err := strconv.Atoi(s[6:])
	if err != nil {
		panic(err)
	}
	h.bid = bid

	return h
}
