/*
Package dataprovider ...
*/
package dataprovider

import (
  "bufio"
  "github.com/alexgunkel/go_algorithms/model"
  "github.com/stretchr/testify/assert"
  "os"
  "strconv"
  "testing"
)

type sorter func(*model.Sortables)

/*
MergeResult ...
*/
var MergeResult model.Sortables

/*
Readlines ...
Readlines reads a whole file into memory
and returns a slice of its lines.
*/
func Readlines(dataFile string) (model.Sortables, error) {
  file, err := os.Open(dataFile)
  if err != nil {
    panic(err)
  }
  defer file.Close()

  var lines model.Sortables
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    var line int64
    line, _ = strconv.ParseInt(scanner.Text(), 10, 64)
    lines = append(lines, line)
  }
  return lines, scanner.Err()
}

/*
Build a sorted list of length integers
and return as model.Sortables
*/
func getSortedList(length int) model.Sortables {
  var sortedListOfIntegers model.Sortables

  for i := 1; i <= length; i++ {
    sortedListOfIntegers = append(sortedListOfIntegers, int64(i))
  }
  return sortedListOfIntegers
}

/*
Build a sorted list of integers descending from
length to one
*/
func getSortedListInverted(length int) model.Sortables {
  var sortedListOfIntegers model.Sortables

  for i := length; i > 0; i-- {
    sortedListOfIntegers = append(sortedListOfIntegers, int64(i))
  }
  return sortedListOfIntegers
}

/*
BenchmarkWithSortedList ...

Helper function to make benchmark testing with sorted lists
easier. Test function with a sorted list (ascending).
*/
func BenchmarkWithSortedList(b *testing.B, f sorter, length int) {
  content := getSortedList(length)

  for n := 0; n < b.N; n++ {
    f(&content)
  }

  MergeResult = content
}

/*
BenchmarkWithSortedListInverted ...

Helper function to make benchmark testing with sorted lists
easier. Test function with a sorted list (descending).
*/
func BenchmarkWithSortedListInverted(b *testing.B, f sorter, length int) {
  content := getSortedListInverted(length)

  for n := 0; n < b.N; n++ {
    f(&content)
  }

  MergeResult = content
}

/*
BenchmarkSort ...
*/
func BenchmarkSort(b *testing.B, f sorter, amount int64) {
  input, _ := Readlines("./dataprovider/data/random" + strconv.FormatInt(amount, 10))
  content := model.Sortables(input)
  orig := content
  for n := 0; n < b.N; n++ {
    content = orig
    f(&content)
  }

  MergeResult = content
}

/*
SorterTest ...

Test sort-function with small test-data-list
*/
func SorterTest(t *testing.T, f sorter) {
  content, _ := Readlines("./dataprovider/data/random10")
  originalLength := len(content)
  f(&content)
  assert.Equal(t, originalLength, len(content))

  for key, number := range content {
    if key > 0 {
      compare := content[key-1]
      assert.Truef(t, compare < number, "Asssert that %d is greater then %d", number, compare)
    }
  }
}
