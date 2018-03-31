/*
Package sort ...
*/
package sort

/*
InsertionSort ...
*/
func InsertionSort(input *sortables) {
	sortables := *input
	*input = insertionSort(sortables)
}

func insertionSort(input sortables) sortables {
	length := len(input)
	for i := 1; i < length; i++ {
		for j := i; j > 0; j-- {
			if input[j] < input[j-1] {
				tmp := input[j]
				input[j] = input[j-1]
				input[j-1] = tmp
			}
		}
	}

	return input
}
