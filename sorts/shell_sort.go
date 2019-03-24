package sorts

import "fmt"

func ShellSort(input []int) {

	//增量初始选长度一般 后面依次减半

	n := len(input)

	swap := func(a, b *int) {
		if *a < *b {
			*a, *b = *b, *a
		}
	}

	for h := len(input) / 2; h > 0; h = h / 2 {

		for i := h; i < n; i++ {
			for j := i; j >= h && input[j] < input[j-h]; j -= h {
				swap(&input[j], &input[j-h])
			}
		}

		fmt.Println(input)
	}
}
