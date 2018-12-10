package mylinkedlist

import (
	"testing"
)

func main() {

}

func TestConstructor(t *testing.T) {
	s := Constructor()

	s.AddAtHead(1)
	s.PrintALL()
	s.AddAtHead(2)
	s.PrintALL()
	s.AddAtTail(4)
	s.PrintALL()
	s.AddAtIndex(2, 3)
	s.PrintALL()
	s.DeleteAtIndex(2)
	s.PrintALL()
	s.DeleteAtIndex(0)
	s.PrintALL()
}

func TestA(t *testing.T) {
	obj := Constructor()

	obj.AddAtHead(1)
	param1 := obj.Get(0)
	print(param1)
	obj.AddAtTail(3)
	obj.AddAtIndex(1, 2)
	obj.DeleteAtIndex(1)
	obj.PrintALL()
}
func TestB(t *testing.T) {
	obj := Constructor()

	obj.AddAtHead(1)
	obj.AddAtIndex(1, 2)
	obj.DeleteAtIndex(1)
	obj.AddAtIndex(1, 9)
	obj.Get(1)
	obj.Get(0)
	obj.Get(2)
	obj.PrintALL()
}

func TestV(t *testing.T) {

	obj := Constructor()
	obj.AddAtHead(5)
	obj.AddAtHead(2)
	obj.DeleteAtIndex(1)
	obj.AddAtIndex(1, 9)
	obj.AddAtHead(4)
	obj.AddAtHead(9)
	obj.AddAtHead(8)
	obj.Get(3)
	obj.AddAtTail(1)
	obj.AddAtIndex(3, 6)
	obj.AddAtHead(3)
}
