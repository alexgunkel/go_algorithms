//shell_sort_test.go
package sort

import (
	"github.com/alexgunkel/go_algorithms/dataprovider"
	"github.com/alexgunkel/go_algorithms/model"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestShellSort(testing *testing.T) {
	content, _ := dataprovider.ReadLines("./data/random10")
	originalLength := len(content)
	ShellSort(&content)
	assert.Equal(testing, originalLength, len(content))

	for key, number := range content {
		if key > 0 {
			compare := content[key-1]
			assert.Truef(testing, compare < number, "Asssert that %d is greater then %d", number, compare)
		}
	}
}

var shellResult model.Sortables

func benchmarkShellSort(amount int64, b *testing.B) {
	file := "./data/random" + strconv.FormatInt(amount, 10)
	content, _ := dataprovider.ReadLines(file)
	var newResult model.Sortables
	for n := 0; n < b.N; n++ {
		newResult = shellSort(content)
	}

	shellResult = newResult
}

func BenchmarkShellSort10(b *testing.B)     { benchmarkShellSort(10, b) }
func BenchmarkShellSort100(b *testing.B)    { benchmarkShellSort(100, b) }
func BenchmarkShellSort1000(b *testing.B)   { benchmarkShellSort(1000, b) }
func BenchmarkShellSort10000(b *testing.B)  { benchmarkShellSort(10000, b) }
func BenchmarkShellSort100000(b *testing.B) { benchmarkShellSort(100000, b) }
