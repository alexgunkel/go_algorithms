package sort

import (
	"github.com/alexgunkel/go_algorithms/dataprovider"
	"github.com/alexgunkel/go_algorithms/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInsertionSort(testing *testing.T) {
	input, _ := dataprovider.ReadLines("./data/random10")
	content := model.Sortables(input)
	InsertionSort(&content)

	for key, number := range content {
		if key > 0 {
			compare := content[key-1]
			assert.Truef(testing, compare < number, "Asssert that %d is greater then %d", number, compare)
		}
	}
}

func BenchmarkInsertionSort10(b *testing.B)     { dataprovider.BenchmarkSort(b, InsertionSort, 10) }
func BenchmarkInsertionSort100(b *testing.B)    { dataprovider.BenchmarkSort(b, InsertionSort, 100) }
func BenchmarkInsertionSort1000(b *testing.B)   { dataprovider.BenchmarkSort(b, InsertionSort, 1000) }
func BenchmarkInsertionSort10000(b *testing.B)  { dataprovider.BenchmarkSort(b, InsertionSort, 10000) }
func BenchmarkInsertionSort100000(b *testing.B) { dataprovider.BenchmarkSort(b, InsertionSort, 100000) }

func BenchmarkSortedInsertion10(b *testing.B) {
	dataprovider.BenchmarkWithSortedList(b, InsertionSort, 10)
}
func BenchmarkSortedInsertion100(b *testing.B) {
	dataprovider.BenchmarkWithSortedList(b, InsertionSort, 100)
}
func BenchmarkSortedInsertion1000(b *testing.B) {
	dataprovider.BenchmarkWithSortedList(b, InsertionSort, 1000)
}
func BenchmarkSortedInsertion10000(b *testing.B) {
	dataprovider.BenchmarkWithSortedList(b, InsertionSort, 10000)
}

func BenchmarkInvertedSortedInsertion10(b *testing.B) {
	dataprovider.BenchmarkWithSortedListInverted(b, InsertionSort, 10)
}
func BenchmarkInvertedSortedInsertion100(b *testing.B) {
	dataprovider.BenchmarkWithSortedListInverted(b, InsertionSort, 100)
}
func BenchmarkInvertedSortedInsertion1000(b *testing.B) {
	dataprovider.BenchmarkWithSortedListInverted(b, InsertionSort, 1000)
}
func BenchmarkInvertedSortedInsertion10000(b *testing.B) {
	dataprovider.BenchmarkWithSortedListInverted(b, InsertionSort, 10000)
}
