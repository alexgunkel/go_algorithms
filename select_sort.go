package sort

/*
SelectSort ...
*/
func SelectSort(input *sortables) {
	content := *input
	var output sortables
	var length int
	var currentLength int
	var index int
	var minimum int64
	length = len(content)
	currentLength = length
	for i := 0; i < length; i++ {
		minimum = content[0]
		index = 0
		for j := 1; j < currentLength; j++ {
			if content[j] < minimum {
				index = j
				minimum = content[j]
			}
		}
		output = append(output, content[index])
		content = append(content[:index], content[index+1:]...)
		currentLength--
	}

	*input = output
}
