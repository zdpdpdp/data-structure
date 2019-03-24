package sorts

func QuickSort(input []int) {
	quickSort(input, 0, len(input)-1)
}
func quickSort(input []int, begin int, end int) {
	if begin < end {
		mid := partition(input, begin, end)
		quickSort(input, begin, mid-1)
		quickSort(input, mid+1, end)
	}
}

func partition(input []int, begin int, end int) int {

	key := input[begin]

	low := begin
	high := end

	for low < high {
		for low < high && input[high] >= key {
			high--
		}
		input[low] = input[high]
		for low < high && input[low] <= key {
			low++
		}
		input[high] = input[low]
	}

	input[low] = key

	return low
}
