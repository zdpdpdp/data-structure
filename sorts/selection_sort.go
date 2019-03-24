package sorts

func SelectionSort(input []int) {

	for i := 0; i < len(input)-1; i++ {
		index := i
		for j := i + 1; j < len(input); j++ {
			if input[j] < input[index] {
				index = j
			}
		}
		if index != i {
			input[index], input[i] = input[i], input[index]
		}
	}
}
