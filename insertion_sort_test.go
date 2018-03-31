package sort

import (
	"./dataprovider"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestInsertionSort(testing *testing.T) {
	input, _ := readLines("./data/random10")
	content := sortables(input)
	InsertionSort(&content)

	for key, number := range content {
		if key > 0 {
			compare := content[key-1]
			assert.Truef(testing, compare < number, "Asssert that %d is greater then %d", number, compare)
		}
	}
}

var insertionResult sortables

func benchmarkInsertionSort(amount int64, b *testing.B) {
	file := "./data/random" + strconv.FormatInt(amount, 10)
	input, _ := readLines(file)
	content := sortables(input)
	var newResult sortables
	for n := 0; n < b.N; n++ {
		newResult = insertionSort(content)
	}

	insertionResult = newResult
}

func BenchmarkInsertionSort10(b *testing.B)     { benchmarkInsertionSort(10, b) }
func BenchmarkInsertionSort100(b *testing.B)    { benchmarkInsertionSort(100, b) }
func BenchmarkInsertionSort1000(b *testing.B)   { benchmarkInsertionSort(1000, b) }
func BenchmarkInsertionSort10000(b *testing.B)  { benchmarkInsertionSort(10000, b) }
func BenchmarkInsertionSort100000(b *testing.B) { benchmarkInsertionSort(100000, b) }

func BenchmarkSortedInsertion1(b *testing.B) {
	dataprovider.BenchmarkWithSortedList(b, InsertionSort, 1)
}
func BenchmarkSortedInsertion2(b *testing.B) {
	dataprovider.BenchmarkWithSortedList(b, InsertionSort, 2)
}
func BenchmarkSortedInsertion4(b *testing.B) {
	dataprovider.BenchmarkWithSortedList(b, InsertionSort, 4)
}
func BenchmarkSortedInsertion8(b *testing.B) {
	dataprovider.BenchmarkWithSortedList(b, InsertionSort, 8)
}

func BenchmarkInvertedSortedInsertion1(b *testing.B) {
	benchmarkWithSortedListInverted(b, InsertionSort, 1)
}
func BenchmarkInvertedSortedInsertion2(b *testing.B) {
	benchmarkWithSortedListInverted(b, InsertionSort, 2)
}
func BenchmarkInvertedSortedInsertion4(b *testing.B) {
	benchmarkWithSortedListInverted(b, InsertionSort, 4)
}
func BenchmarkInvertedSortedInsertion8(b *testing.B) {
	benchmarkWithSortedListInverted(b, InsertionSort, 8)
}
