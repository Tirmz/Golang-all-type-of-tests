package sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBubbleSortAsc(t *testing.T) {
	//Init
	elements := GetElements(10)

	assert.NotNil(t, elements, "Array is nil")
	assert.EqualValues(t, 10, len(elements), "Wrong number of elements")
	assert.EqualValues(t, 9, elements[0], "First element should be 9")
	assert.EqualValues(t, 0, elements[len(elements)-1], "Last element should be 0")

	//Execution
	BubbleSort(elements)

	//Validation
	assert.NotNil(t, elements, "Array is nil")
	assert.EqualValues(t, 10, len(elements), "Wrong number of elements")
	assert.EqualValues(t, 0, elements[0], "First element should be 0")
	assert.EqualValues(t, 9, elements[len(elements)-1], "Last element should be 9")
}

func TestSortAsc(t *testing.T) {
	//Init
	elements := GetElements(10)

	//Execution
	Sort(elements)

	//Validation
	assert.EqualValues(t, 0, elements[0], "First element should be 0")
	assert.EqualValues(t, 9, elements[len(elements)-1], "Last element should be 9")
}

func BenchmarkBubbleSort(b *testing.B) {
	elements := GetElements(10000)

	for i := 0; i < b.N; i++ {
		BubbleSort(elements)
	}
}

func BenchmarkSort(b *testing.B) {
	elements := GetElements(10000)

	for i := 0; i < b.N; i++ {
		Sort(elements)
	}
}
