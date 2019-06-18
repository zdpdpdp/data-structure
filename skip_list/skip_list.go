package skip_list

import (
	"fmt"
	"math/rand"
	"time"
)

//MaxLevel 限定最大层数
const MaxLevel = 16

type SkipList struct {
	maxLevel     int
	currentLevel int
	head         *Node
	size         int
}

type Node struct {
	Val   int
	level int //跨越层数
	next  []*Node
}

//New 初始化跳跃表
func New() *SkipList {

	node := new(Node)

	node.next = make([]*Node, MaxLevel)

	return &SkipList{
		maxLevel: MaxLevel,
		head:     node,
	}
}

//randomLevel 采用抛硬币的方法, 如果硬币朝上, 则再抛一次, 直到硬币朝下为止
func (s *SkipList) randomLevel() int {

	var level int

	for level < s.maxLevel {
		level++
		rand.Seed(time.Now().UnixNano())
		if rand.Int()%2 == 1 {
			break
		}
	}

	return level

}

//Insert 插入
func (s *SkipList) Insert(val int) {
	level := s.randomLevel()
	node := &Node{
		level: level,
		Val:   val,
		next:  make([]*Node, level),
	}

	//定位插入的位置
	var preNodes = make([]*Node, level)

	var temp = s.head

	for i := level - 1; i >= 0; i-- {
		for temp.next[i] != nil && temp.next[i].Val < val { //顶层 自上而下查找
			temp = temp.next[i]
		}

		//每一次循环后, 都能查找到插入位置的前驱节点
		preNodes[i] = temp
	}

	for i, preNode := range preNodes {
		node.next[i] = preNode.next[i]
		preNode.next[i] = node
	}

	s.size++ //记录节点个数

	if s.currentLevel < level { //记录最大层数
		s.currentLevel = level
	}

}

//Delete 删除节点, 为了方便, 不考虑重复节点
func (s *SkipList) Delete(val int) {

	var temp = s.head
	var flag bool
	for i := s.currentLevel; i >= 0; i-- {
		for temp.next[i] != nil && temp.next[i].Val < val {
			temp = temp.next[i]
		}

		if temp.next[i] != nil && temp.next[i].Val == val { //下一个节点的值等于删除值
			flag = true
			temp.next[i] = temp.next[i].next[i]
		}
	}

	if flag {
		s.size--
	}
}

func (s *SkipList) Len() int {
	return s.size
}

//Find 查找
func (s *SkipList) Find(val int) bool {
	tmp := s.head

	for i := s.currentLevel - 1; i >= 0; i-- {
		for tmp.next[i] != nil && tmp.next[i].Val < val {
			tmp = tmp.next[i]
		}
	}

	return tmp.next[0] != nil && tmp.next[0].Val == val
}

//Print 打印跳跃表
func (s *SkipList) Print() {

	for i := s.currentLevel - 1; i >= 0; i-- {
		temp := s.head
		for temp.next[i] != nil {
			fmt.Print(temp.next[i].Val, "\t")
			temp = temp.next[i]
		}
		fmt.Println()
	}
}
