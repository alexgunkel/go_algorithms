/*
Package sort ...
*/
package sort

import (
	"fmt"

	"github.com/alexgunkel/go_algorithms/model"
)

/*
QuickSort ...
*/
func QuickSort(input *model.Sortables) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Error while partitioning", err)
		}
	}()

	data := *input

	quickSort(data[0:])

	*input = data
}

func quickSort(input model.Sortables) {
	if len(input) < 2 {
		return
	}
	i := partition(input)
	quickSort(input[:i])
	quickSort(input[i+1:])
}

func partition(input model.Sortables) int {
	high := len(input) - 1
	i := 1
	j := high
	value := input[0]

	for {
		for input[i] <= value && i != high {
			i++
		}

		for input[j] > value && j != 0 {
			j--
		}

		if i >= j {
			break
		}

		input[i], input[j] = input[j], input[i]
	}

	input[0], input[j] = input[j], input[0]

	return j
}
