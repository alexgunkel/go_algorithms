// select_sort_test.go
package sort

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestSelectSort(testing *testing.T) {
	input, _ := readLines("./data/random10")
	content := sortables(input)
	SelectSort(&content)

	for key, number := range content {
		if key > 0 {
			compare := content[key-1]
			assert.Truef(testing, compare < number, "Asssert that %d is greater then %d", number, compare)
		}
	}
}

var selectResult sortables

func benchmarkSelectSort(amount int64, b *testing.B) {
	file := "./data/random" + strconv.FormatInt(amount, 10)
	input, _ := readLines(file)
	content := sortables(input)

	for n := 0; n < b.N; n++ {
		SelectSort(&content)
	}

	selectResult = content
}

func BenchmarkSelectSort10(b *testing.B)     { benchmarkMergeSort(10, b) }
func BenchmarkSelectSort100(b *testing.B)    { benchmarkMergeSort(100, b) }
func BenchmarkSelectSort1000(b *testing.B)   { benchmarkMergeSort(1000, b) }
func BenchmarkSelectSort10000(b *testing.B)  { benchmarkMergeSort(10000, b) }
func BenchmarkSelectSort100000(b *testing.B) { benchmarkMergeSort(100000, b) }
