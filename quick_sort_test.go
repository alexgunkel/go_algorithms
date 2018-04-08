package sort

import (
	"testing"

	"github.com/alexgunkel/go_algorithms/dataprovider"
	"github.com/alexgunkel/go_algorithms/model"
	"github.com/stretchr/testify/assert"
)

func TestPartition(t *testing.T) {
	for key, testDatum := range partitioningTestData() {
		sortables := testDatum.input
		result := partition(sortables)
		assert.Truef(t, result == testDatum.result, "Assert that partition finds correct position (%d) but got %d in %d", testDatum.result, result, key)
		for i := 0; i < len(sortables); i++ {
			if i < result {
				assert.Truef(t, sortables[i] <= sortables[result], "Value at position %d should be smaller or equal %d. Test %d", i, sortables[i], sortables[result], key)
			}

			if i > result {
				assert.Truef(t, sortables[i] > sortables[result], "Value at position %d should be larger than %d. Test %d", i, sortables[i], sortables[result], key)
			}
		}
	}
}

type partitioningTestDatum struct {
	input  model.Sortables
	result int
}

func partitioningTestData() []partitioningTestDatum {
	return []partitioningTestDatum{
		{input: model.Sortables{3, 4}, result: 0},
		{input: model.Sortables{3, 2}, result: 1},
		{input: model.Sortables{1, 2}, result: 0},
		{input: model.Sortables{3, 1, 2}, result: 2},
		{input: model.Sortables{3, 1, 2, 4, 5}, result: 2},
		{input: model.Sortables{3, 1, 2, 1, 5, 8}, result: 3},
		{input: model.Sortables{3, 1, 2, 3, 4, 5, 6, 7, 8}, result: 3},
	}
}

func TestQuickSort(t *testing.T) {
	dataprovider.SorterTest(t, QuickSort)
}

func BenchmarkQuickSort10(b *testing.B)     { dataprovider.BenchmarkSort(b, QuickSort, 10) }
func BenchmarkQuickSort100(b *testing.B)    { dataprovider.BenchmarkSort(b, QuickSort, 100) }
func BenchmarkQuickSort1000(b *testing.B)   { dataprovider.BenchmarkSort(b, QuickSort, 1000) }
func BenchmarkQuickSort10000(b *testing.B)  { dataprovider.BenchmarkSort(b, QuickSort, 10000) }
func BenchmarkQuickSort100000(b *testing.B) { dataprovider.BenchmarkSort(b, QuickSort, 100000) }
