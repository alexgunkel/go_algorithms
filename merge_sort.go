/*
Package sort

Merge-sort algorithm

Author Alexander Gunkel <alexandergunkel@gmail.com>
*/

package sort

/*
MergeSort ...

@param	*sortables listOfThings	pointer to a slice of things to sort
*/
func MergeSort(listOfThings *sortables) {
	input := *listOfThings
	*listOfThings = mergeSort(input)
}

/*
MergeSortInline ...

@param	*sortables listOfThings	pointer to a slice of things to sort
*/
func MergeSortInline(listOfThings *sortables) {
	input := *listOfThings
	lastOffset := len(input) - 1
	mergeSortInline(listOfThings, 0, lastOffset)
}

/*
MergeSortMixed ...

@param	*sortables listOfThings	pointer to a slice of things to sort
*/
func MergeSortMixed(listOfThings *sortables) {
	input := *listOfThings
	*listOfThings = mergeSortMixed(input)
}

/*
Backend function. Combines merge-sort with insertion-sort
for small lists and go-routines for very large lists.

This turns out to be quite fast on long lists while being
slow on short lists.
*/
func mergeSort(input sortables) sortables {
	var length, middle int
	length = len(input)
	if length <= 1 {
		return input
	}

	middle = length / 2

	if length < 5000 {
		return merge(mergeSortMixed(input[0:middle]), mergeSortMixed(input[middle:length]))
	}

	c1 := make(chan sortables)

	go func() {
		c1 <- mergeSort(input[0:middle])
	}()

	go func() {
		c1 <- mergeSort(input[middle:length])
	}()

	return merge(<-c1, <-c1)
}

/*
Mixed merge-sort backend-function
*/
func mergeSortMixed(input sortables) sortables {
	var length, middle int
	length = len(input)
	if length < 50 {
		return insertionSort(input)
	}

	middle = length / 2

	return merge(mergeSortMixed(input[0:middle]), mergeSortMixed(input[middle:length]))
}

func merge(partOne sortables, partTwo sortables) (result sortables) {
	lengthPartOne := len(partOne)
	lengthPartTwo := len(partTwo)
	positionPartOne := 0
	positionPartTwo := 0
	total := lengthPartOne + lengthPartTwo
	for i := 0; i < total; i++ {
		if lengthPartTwo == 0 || (lengthPartOne != 0 && partOne[positionPartOne] < partTwo[positionPartTwo]) {
			result = append(result, partOne[positionPartOne])
			positionPartOne++
			lengthPartOne--
		} else {
			result = append(result, partTwo[positionPartTwo])
			positionPartTwo++
			lengthPartTwo--
		}
	}

	return result
}

// Use only one slice. Don't copy arrays
//
// Function expects slice as pointer to avoid any copying
func mergeSortInline(input *sortables, start int, end int) {
	var middle int
	diff := end - start

	if diff < 1 {
		return
	}

	middle = (diff+1)/2 + start
	mergeSortInline(input, start, middle-1)
	mergeSortInline(input, middle, end)
	mergeParts(input, start, middle, end)
}

// Merge parts of arrays
//
// Used to merge on one and the same array
func mergeParts(input *sortables, startFirst int, startSecond int, end int) {
	listOfThings := *input
	var tmp int64
	offsetFirst := startFirst
	offsetSecond := startSecond

	for offsetFirst < offsetSecond && offsetSecond <= end {
		if listOfThings[offsetFirst] > listOfThings[offsetSecond] {
			tmp = listOfThings[offsetSecond]
			for i := offsetSecond; i > offsetFirst; i-- {
				listOfThings[i] = listOfThings[i-1]
			}
			listOfThings[offsetFirst] = tmp
			offsetSecond++
			offsetFirst++
		} else {
			offsetFirst++
		}
	}

	*input = listOfThings
}
