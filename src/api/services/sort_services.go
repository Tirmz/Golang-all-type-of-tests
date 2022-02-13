package services

import (
	"Golang-all-type-of-tests/src/utils/sort"
)

func Sort(elements []int) {
	if len(elements) < 20000 {
		sort.Sort(elements)
	} else {
		sort.BubbleSort(elements)
	}

}
