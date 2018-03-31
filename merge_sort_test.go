// merge_sort_test.go
package sort

import (
	"./dataprovider"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestMergeSort(testing *testing.T) {
	input, _ := readLines("./data/random10")
	content := sortables(input)
	originalLength := len(content)
	MergeSort(&content)
	assert.Equal(testing, originalLength, len(content))

	for key, number := range content {
		if key > 0 {
			compare := content[key-1]
			assert.Truef(testing, compare < number, "Asssert that %d is greater then %d", number, compare)
		}
	}
}

func TestMergeSortInline(testing *testing.T) {
	input, _ := readLines("./data/random10")
	content := sortables(input)
	originalLength := len(content)
	MergeSortInline(&content)
	assert.Equal(testing, originalLength, len(content))

	for key, number := range content {
		if key > 0 {
			compare := content[key-1]
			assert.Truef(testing, compare < number, "Asssert that %d is greater then %d", number, compare)
		}
	}
}

func TestMergeParts(testing *testing.T) {
	testInput := sortables{1, 3, 5, 2, 4, 6, 1, 2}
	expected := sortables{1, 2, 3, 4, 5, 6, 1, 2}
	mergeParts(&testInput, 0, 3, 6)

	assert.Equal(testing, expected, testInput)
}

func benchmarkMergeSort(amount int64, b *testing.B) {
	file := "./data/random" + strconv.FormatInt(amount, 10)
	input, _ := readLines(file)
	content := sortables(input)

	for n := 0; n < b.N; n++ {
		MergeSort(&content)
	}

	mergeResult = content
}

func BenchmarkMergeSort10(b *testing.B)     { benchmarkMergeSort(10, b) }
func BenchmarkMergeSort100(b *testing.B)    { benchmarkMergeSort(100, b) }
func BenchmarkMergeSort1000(b *testing.B)   { benchmarkMergeSort(1000, b) }
func BenchmarkMergeSort10000(b *testing.B)  { benchmarkMergeSort(10000, b) }
func BenchmarkMergeSort100000(b *testing.B) { benchmarkMergeSort(100000, b) }

func benchmarkMixedMergeSort(amount int64, b *testing.B) {
	file := "./data/random" + strconv.FormatInt(amount, 10)
	input, _ := readLines(file)
	content := sortables(input)

	for n := 0; n < b.N; n++ {
		MergeSortMixed(&content)
	}

	mergeResult = content
}

func BenchmarkMixedMergeSort10(b *testing.B)     { benchmarkMixedMergeSort(10, b) }
func BenchmarkMixedMergeSort100(b *testing.B)    { benchmarkMixedMergeSort(100, b) }
func BenchmarkMixedMergeSort1000(b *testing.B)   { benchmarkMixedMergeSort(1000, b) }
func BenchmarkMixedMergeSort10000(b *testing.B)  { benchmarkMixedMergeSort(10000, b) }
func BenchmarkMixedMergeSort100000(b *testing.B) { benchmarkMixedMergeSort(100000, b) }

func benchmarkInlineMergeSort(amount int64, b *testing.B) {
	file := "./data/random" + strconv.FormatInt(amount, 10)
	input, _ := readLines(file)
	content := sortables(input)

	for n := 0; n < b.N; n++ {
		MergeSortInline(&content)
	}

	mergeResult = content
}

func BenchmarkInlineMergeSort10(b *testing.B)     { benchmarkMixedMergeSort(10, b) }
func BenchmarkInlineMergeSort100(b *testing.B)    { benchmarkMixedMergeSort(100, b) }
func BenchmarkInlineMergeSort1000(b *testing.B)   { benchmarkMixedMergeSort(1000, b) }
func BenchmarkInlineMergeSort10000(b *testing.B)  { benchmarkMixedMergeSort(10000, b) }
func BenchmarkInlineMergeSort100000(b *testing.B) { benchmarkMixedMergeSort(100000, b) }

func BenchmarkSortedMerge1(b *testing.B) { benchmarkWithSortedList(b, MergeSort, 1) }
func BenchmarkSortedMerge2(b *testing.B) { benchmarkWithSortedList(b, MergeSort, 2) }
func BenchmarkSortedMerge4(b *testing.B) { benchmarkWithSortedList(b, MergeSort, 4) }
func BenchmarkSortedMerge8(b *testing.B) { benchmarkWithSortedList(b, MergeSort, 8) }

func BenchmarkSortedMergeInline1(b *testing.B) { benchmarkWithSortedList(b, MergeSortInline, 1) }
func BenchmarkSortedMergeInline2(b *testing.B) { benchmarkWithSortedList(b, MergeSortInline, 2) }
func BenchmarkSortedMergeInline4(b *testing.B) { benchmarkWithSortedList(b, MergeSortInline, 4) }
func BenchmarkSortedMergeInline8(b *testing.B) { benchmarkWithSortedList(b, MergeSortInline, 8) }

func BenchmarkSortedMergeMixed1(b *testing.B) { benchmarkWithSortedList(b, MergeSortMixed, 1) }
func BenchmarkSortedMergeMixed2(b *testing.B) { benchmarkWithSortedList(b, MergeSortMixed, 2) }
func BenchmarkSortedMergeMixed4(b *testing.B) { benchmarkWithSortedList(b, MergeSortMixed, 4) }
func BenchmarkSortedMergeMixed8(b *testing.B) { benchmarkWithSortedList(b, MergeSortMixed, 8) }

func BenchmarkSortedMerge1Inverted(b *testing.B) {
	benchmarkWithSortedListInverted(b, MergeSort, 1)
}
func BenchmarkSortedMerge2Inverted(b *testing.B) {
	benchmarkWithSortedListInverted(b, MergeSort, 2)
}
func BenchmarkSortedMerge4Inverted(b *testing.B) {
	benchmarkWithSortedListInverted(b, MergeSort, 4)
}
func BenchmarkSortedMerge8Inverted(b *testing.B) {
	benchmarkWithSortedListInverted(b, MergeSort, 8)
}

func BenchmarkInvertedMergeInline1(b *testing.B) {
	benchmarkWithSortedListInverted(b, MergeSortInline, 1)
}
func BenchmarkInvertedMergeInline2(b *testing.B) {
	benchmarkWithSortedListInverted(b, MergeSortInline, 2)
}
func BenchmarkInvertedMergeInline4(b *testing.B) {
	benchmarkWithSortedListInverted(b, MergeSortInline, 4)
}
func BenchmarkInvertedMergeInline8(b *testing.B) {
	benchmarkWithSortedListInverted(b, MergeSortInline, 8)
}
func BenchmarkInvertedMergeMixed1(b *testing.B) {
	benchmarkWithSortedListInverted(b, MergeSortMixed, 1)
}
func BenchmarkInvertedMergeMixed2(b *testing.B) {
	benchmarkWithSortedListInverted(b, MergeSortMixed, 2)
}
func BenchmarkInvertedMergeMixed4(b *testing.B) {
	benchmarkWithSortedListInverted(b, MergeSortMixed, 4)
}
func BenchmarkInvertedMergeMixed8(b *testing.B) {
	benchmarkWithSortedListInverted(b, MergeSortMixed, 8)
}
