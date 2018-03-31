/*
Package sort ...

Shell-sort algorithm

Author Alexander Gunkel <alexandergunkel@gmail.com>
*/
package sort

import "github.com/alexgunkel/go_algorithms/model"

/*
ShellSort ...
*/
func ShellSort(input *model.Sortables) {
	args := *input
	*input = shellSort(args)
}

func shellSort(input model.Sortables) model.Sortables {
	var length, interval int
	length = len(input)
	interval = 1

	for interval < length/3 {
		interval = interval*3 + 1
	}

	for interval > 0 {
		for i := interval; i < length; i++ {
			for j := i; j >= interval && input[j-interval] > input[j]; j = j - interval {
				changePositions(&input, j, j-interval)
			}
		}
		interval = interval / 3
	}

	return input
}

func changePositions(slPtr *model.Sortables, posOne int, posTwo int) {
	sl := *slPtr
	var tmp int64
	tmp = sl[posOne]
	sl[posOne] = sl[posTwo]
	sl[posTwo] = tmp
	*slPtr = sl
}
