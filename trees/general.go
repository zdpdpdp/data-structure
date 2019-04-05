package trees

//中序遍历
func InOrderTraverse(node *Node, res *[]int) {
	if node != nil {
		*res = append(*res, node.Val)
	}
	InOrderTraverse(node.Left, res)
	InOrderTraverse(node.Right, res)
}

//先序遍历
func FirstOrderTraverse(node *Node, res *[]int) {
	if node != nil {
		FirstOrderTraverse(node.Left, res)
		*res = append(*res, node.Val)
		FirstOrderTraverse(node.Right, res)
	}
}

//后续遍历
func LastOrderTraverse(node *Node, res *[]int) {
	if node == nil {
		LastOrderTraverse(node.Left, res)
		LastOrderTraverse(node.Right, res)
		*res = append(*res, node.Val)
	}
}
func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//层次遍历
func LevelTraverse(node *Node, res *[]int) {
	if node == nil {
		return
	}

	queue := NewQueue()
	queue.Push(node)
	for queue.len() > 0 {
		popNode, _ := queue.Pop()
		*res = append(*res, popNode.Val)
		if popNode.Left != nil {
			queue.Push(popNode.Left)
		}
		if popNode.Right != nil {
			queue.Push(popNode.Right)
		}
	}

}

//求深度
func Depth(node *Node) int {
	if node == nil {
		return 0
	}

	return maxInt(Depth(node.Left)+1, Depth(node.Right)+1)
}

func Build(source []int) *Node {
	//类似构造一个满二叉树的形式
	return createTree(0, source)
}

func createTree(n int, source []int) *Node {
	if len(source) < n+1 {
		return nil
	}

	if source[n] == -1 {
		return nil
	}
	res := &Node{Val: source[n]}
	res.Left = createTree(2*n+1, source)
	res.Right = createTree(2*n+2, source)
	return res

}

////打印一颗二叉树
//func PrintTraverse(node *Node) {
//	depth := Depth(node)
//}
