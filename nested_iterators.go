package main

import (
	"fmt"
	"iter"
)

// Iterator that returns numbers from start to end
func rangeIterator(start, end int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := start; i <= end; i++ {
			if !yield(i) {
				return
			}
		}
	}
}

// Iterator that returns other iterators
func nestedIterator(start, end, size int) iter.Seq[iter.Seq[int]] {
	return func(yield func(iter.Seq[int]) bool) {
		for i := start; i <= end; i += size {
			nestedEnd := i + size - 1
			if nestedEnd > end {
				nestedEnd = end
			}

			// rangeIterator is another, similar iterator func
			subRange := rangeIterator(i, nestedEnd) // HLxxx

			if !yield(subRange) {
				return
			}
		}
	}
}

func main() {
	chunkedRanges := nestedIterator(1, 10, 3)

	for chunk := range chunkedRanges {
		fmt.Println("New chunk:")

		// Iterate through the inner iterator
		for num := range chunk {
			fmt.Printf("  %d\n", num)
		}
	}
}
