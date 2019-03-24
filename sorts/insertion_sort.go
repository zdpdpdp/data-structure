package sorts

func InsertionSort(input []int) { //简单插入排序

	for i := 1; i < len(input); i++ {
		for j := i - 1; j >= 0; j-- {
			if input[j+1] < input[j] { //交换
				input[j], input[j+1] = input[j+1], input[j]
			} else {
				break
			}
		}
	}
}
