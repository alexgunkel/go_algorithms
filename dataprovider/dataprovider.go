/*
Package dataprovider ...
*/
package dataprovider

import (
  "bufio"
  "os"
  "strconv"
  "testing"
)

type sorter func(*[]int64)

var mergeResult []int64

/*
ReadLines ...
readLines reads a whole file into memory
and returns a slice of its lines.
*/
func ReadLines(path string) ([]int64, error) {
  file, err := os.Open(path)
  if err != nil {
    return nil, err
  }
  defer file.Close()

  var lines []int64
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
and return as []int64
*/
func getSortedList(length int) []int64 {
  var sortedListOfIntegers []int64

  for i := 1; i <= 1000*length; i++ {
    sortedListOfIntegers = append(sortedListOfIntegers, int64(i))
  }
  return sortedListOfIntegers
}

/*
Build a sorted list of integers descending from
length to one
*/
func getSortedListInverted(length int) []int64 {
  var sortedListOfIntegers []int64

  for i := 1000 * length; i > 0; i-- {
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

  mergeResult = content
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

  mergeResult = content
}
