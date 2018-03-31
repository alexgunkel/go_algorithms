//shell_sort_test.go
package sort

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestShellSort(testing *testing.T) {
	content, _ := readLines("./data/random10")
	original_length := len(content)
	ShellSort(&content)
	assert.Equal(testing, original_length, len(content))

	for key, number := range content {
		if key > 0 {
			compare := content[key-1]
			assert.Truef(testing, compare < number, "Asssert that %d is greater then %d", number, compare)
		}
	}
}

var shell_result sortables

func benchmarkShellSort(amount int64, b *testing.B) {
	file := "./data/random" + strconv.FormatInt(amount, 10)
	content, _ := readLines(file)
	var new_result sortables
	for n := 0; n < b.N; n++ {
		new_result = shellSort(content)
	}

	shell_result = new_result
}

func BenchmarkShellSort10(b *testing.B)     { benchmarkShellSort(10, b) }
func BenchmarkShellSort100(b *testing.B)    { benchmarkShellSort(100, b) }
func BenchmarkShellSort1000(b *testing.B)   { benchmarkShellSort(1000, b) }
func BenchmarkShellSort10000(b *testing.B)  { benchmarkShellSort(10000, b) }
func BenchmarkShellSort100000(b *testing.B) { benchmarkShellSort(100000, b) }
