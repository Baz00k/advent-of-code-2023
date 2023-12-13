package day05

import (
	"sort"
	"strconv"
	"strings"
)

type valuesRange struct {
	start int
	end   int
}

func mergeOverlapingValueRanges(ranges []*valuesRange) []*valuesRange {
	// Sort the intervals by start value
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].start < ranges[j].start
	})

	// Create a new slice to store the merged intervals
	merged := make([]*valuesRange, 0)

	// Add the first interval to the merged slice
	merged = append(merged, ranges[0])

	// Merge the remaining intervals
	for _, r := range ranges[1:] {
		last := merged[len(merged)-1]
		if r.start <= last.end {
			last.end = max(last.end, r.end)
		} else {
			merged = append(merged, r)
		}
	}

	return merged
}

type remapRange struct {
	start int
	end   int
	shift int
}

func (r remapRange) inRange(vr valuesRange) bool {
	return vr.start >= r.start && vr.end <= r.end
}

func (r remapRange) inRangePartially(vr valuesRange) bool {
	return vr.start <= r.end && vr.end >= r.start
}

func (r remapRange) remap(vr *valuesRange) {
	vr.start += r.shift
	vr.end += r.shift
}

func mergeOverlapingRanges(ranges *[]*remapRange) {
	// We can assume that the ranges are sorted by start value
	rangesMerged := false

	for i := 0; i < len(*ranges)-1; i++ {
		if (*ranges)[i].end >= (*ranges)[i+1].start-1 && (*ranges)[i].shift == (*ranges)[i+1].shift {
			// We can merge the two ranges
			(*ranges)[i].end = (*ranges)[i+1].end
			*ranges = append((*ranges)[:i+1], (*ranges)[i+2:]...)

			rangesMerged = true
		}
	}

	if !rangesMerged {
		return
	}

	mergeOverlapingRanges(ranges)
}

func parseRemapSource(line string) []int {
	// Line should be in the format:
	// 50 98 2
	// Where 50 is destination range start, 98 is source range start, and 2 is range length
	remapSourceParts := strings.Split(strings.TrimSpace(line), " ")
	if len(remapSourceParts) != 3 {
		panic("Invalid remap source")
	}

	remapSource := make([]int, 3)
	for i, part := range remapSourceParts {
		num, err := strconv.Atoi(part)
		if err != nil {
			panic(err)
		}
		remapSource[i] = num
	}

	return remapSource
}

func remapRangeFromLine(line string) *remapRange {
	remapSource := parseRemapSource(line)

	return &remapRange{
		start: remapSource[1],
		end:   remapSource[1] + remapSource[2] - 1,
		shift: remapSource[0] - remapSource[1],
	}
}

// We need to follow a chain of remaps to get the final value
// The chain can be represented as a linked list
// We can then iterate over the list to get the final value
type remapChain struct {
	remaps []*remapRange
	next   *remapChain
}

func (r *remapChain) sortRemaps() {
	sort.Slice(r.remaps, func(i, j int) bool {
		return r.remaps[i].start < r.remaps[j].start
	})

	mergeOverlapingRanges(&r.remaps)

	if r.next != nil {
		r.next.sortRemaps()
	}
}

func (r *remapChain) addRemap(remap *remapRange) {
	r.remaps = append(r.remaps, remap)
}

func (r remapChain) remap(vr valuesRange) []valuesRange {
	// We need to check if our whole range is in the remap range
	// If it is, we can just shift the whole range
	// If it isn't, we need to split the range and shift each part

	// We can assume that the remaps are sorted
	// We can also assume that the remaps don't overlap

	// If the range doesn't fit in any remap, even partially, we can just return the original range
	noMatch := true
	for _, remap := range r.remaps {
		if remap.inRangePartially(vr) {
			noMatch = false
			if remap.inRange(vr) {
				remap.remap(&vr)

				if r.next != nil {
					return r.next.remap(vr)
				}

				return []valuesRange{vr}
			}
		}
	}

	if noMatch {
		if r.next != nil {
			return r.next.remap(vr)
		}

		return []valuesRange{vr}
	}

	// We need to split the range
	var breakPoints []int
	for _, remap := range r.remaps {
		if remap.inRangePartially(vr) {
			breakPoints = append(breakPoints, remap.start, remap.end)
		}
	}
	breakPoints = append(breakPoints, vr.start, vr.end)
	sort.Ints(breakPoints)

	var newRanges []*valuesRange
	rangeStart := vr.start
	for _, breakPoint := range breakPoints {
		if breakPoint <= rangeStart {
			continue
		}

		if breakPoint >= vr.end {
			newRanges = append(newRanges, &valuesRange{
				start: rangeStart,
				end:   vr.end,
			})

			break
		}

		newRanges = append(newRanges, &valuesRange{
			start: rangeStart,
			end:   breakPoint,
		})

		rangeStart = breakPoint + 1
	}

	for _, newRange := range newRanges {
		for _, remap := range r.remaps {
			if remap.inRange(*newRange) {
				remap.remap(newRange)

				break
			}
		}
	}

	newRanges = mergeOverlapingValueRanges(newRanges)

	remappedNewRanges := []valuesRange{}
	for _, newRange := range newRanges {
		if r.next != nil {
			remappedNewRanges = append(remappedNewRanges, r.next.remap(*newRange)...)
		} else {
			remappedNewRanges = append(remappedNewRanges, *newRange)
		}
	}

	return remappedNewRanges
}
