//shell_sort_test.go
package sort

import (
	"github.com/alexgunkel/go_algorithms/dataprovider"
	"testing"
)

func TestShellSort(testing *testing.T) {
	dataprovider.SorterTest(testing, ShellSort)
}

func BenchmarkShellSort10(b *testing.B)     { dataprovider.BenchmarkSort(b, ShellSort, 10) }
func BenchmarkShellSort100(b *testing.B)    { dataprovider.BenchmarkSort(b, ShellSort, 100) }
func BenchmarkShellSort1000(b *testing.B)   { dataprovider.BenchmarkSort(b, ShellSort, 1000) }
func BenchmarkShellSort10000(b *testing.B)  { dataprovider.BenchmarkSort(b, ShellSort, 10000) }
func BenchmarkShellSort100000(b *testing.B) { dataprovider.BenchmarkSort(b, ShellSort, 100000) }
