package trees

import (
	"fmt"
	"testing"
)

func TestBuild(t *testing.T) {
	var data = []int{
		1, 2, 3, 4, 5, 6, 7,
	}

	node := Build(data)
	var order []int
	FirstOrderTraverse(node, &order)
	fmt.Println(order)
}
