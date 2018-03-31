/*
Package sort ...
*/
package sort

import "github.com/alexgunkel/go_algorithms/model"

/*
InsertionSort ...
*/
func InsertionSort(input *model.Sortables) {
	listOfThings := *input
	*input = insertionSort(listOfThings)
}

func insertionSort(input model.Sortables) model.Sortables {
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
