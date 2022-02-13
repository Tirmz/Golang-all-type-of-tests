package services

import (
	"Golang-all-type-of-tests/src/utils/sort"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSort(t *testing.T) {
	elements := sort.GetElements(10)
	Sort(elements)
	assert.NotNil(t, elements, "Array is nil")
	assert.EqualValues(t, 0, elements[0], "First element should be 0")
	assert.EqualValues(t, 9, elements[len(elements)-1], "Last element should be 9")
}

func TestSortMoreThan20000(t *testing.T) {
	elements := sort.GetElements(20001)

	Sort(elements)

	if elements[0] != 0 {
		t.Error("First element should be 0")
	}
	if elements[len(elements)-1] != 20000 {
		t.Error("Last element should be 20000")
	}
}

func BenchmarkBubbleSort10K(b *testing.B) {
	elements := sort.GetElements(10000)

	for i := 0; i < b.N; i++ {
		Sort(elements)
	}
}

func BenchmarkBubbleSort100K(b *testing.B) {
	elements := sort.GetElements(10000)

	for i := 0; i < b.N; i++ {
		Sort(elements)
	}
}
