package skip_list

import (
	"fmt"
	"testing"
)

var l *SkipList

func init() {
	l = New()
}

func TestSkipList(t *testing.T) {
	for i := 0; i < 10; i++ {
		l.Insert(i)
	}

	fmt.Println("first print")
	l.Print()

	for i := 0; i < 10; i++ {
		if !l.Find(i) {
			t.Errorf("cant not find %d", i)
		}
	}

	if l.Len() != 10 {
		t.Errorf("wrong len")
	}
	for i := 11; i < 20; i++ {
		if l.Find(i) {
			t.Errorf("find %d error", i)
		}
	}

	l.Delete(5)
	if l.Find(5) {
		t.Errorf("delete but find %d error", 5)
	}

	l.Delete(5)

	if l.Len() != 9 {
		t.Errorf("wrong len")
	}

	fmt.Println("second print")
	l.Print()
}
