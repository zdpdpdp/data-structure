package sorts

func HeapSort(input []int) {
	//堆排序利用了完全二叉树的性质,建立最大堆最小堆
	createHeap(input)

	l := len(input)
	for i := 0; i < l; i++ {
		//从最大堆里取一个元素,和最后元素交换
		input[0], input[l-1-i] = input[l-1-i], input[0]
		heapAdjust(input[:l-1-i])
	}
}

func heapAdjust(input []int) {

	pos := 0

	length := len(input)
	for child := pos*2 + 1; child < length; child = pos*2 + 1 {

		if child+1 < length && input[child+1] > input[child] {
			child++
		}

		if input[pos] >= input[child] {
			return
		}

		input[pos], input[child] = input[child], input[pos]
		pos = child
	}

}

//建初堆
func createHeap(input []int) {

	//由于完全二叉树 > n/2-1 的节点都是叶子节点,所以调整前面的节点即可
	n := len(input)
	for i := n/2 - 1; i >= 0; i-- {
		heapAdjust(input[i:])
	}
}
