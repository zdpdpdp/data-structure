package trees

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

type Tree interface {
	Insert(int)
	Search(int)
	Remove(int)
}
