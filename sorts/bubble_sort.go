package sorts

func BubbleSort(input []int) {

	for i := 0; i < len(input)-1; i++ {
		var flag = false
		for j := 0; j < len(input)-1-i; j++ {
			if input[j] > input[j+1] {
				input[j], input[j+1] = input[j+1], input[j]
				flag = true
			}
		}
		if flag == false { //没有任何交换
			break
		}
	}
}
