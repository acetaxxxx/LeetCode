package mylinkedlist

type MyNode struct {
	Next *MyNode
	Prev *MyNode
	Val  int
}

func (this *MyNode) Push(val int) {
	n := MyNode{
		Next: nil,
		Prev: nil,
		Val:  val,
	}

	if this.Next != nil {
		n.Next = this.Next
		this.Next = &n
		n.Prev = this
	} else {
		n.Prev = this
		this.Next = &n
	}
}

func (this *MyNode) PopNext() {
	if this.Next == nil{
		return
	}
	if this.Next.Next != nil  {
		this.Next = this.Next.Next
		this.Next.Prev = this
		
		return
	} else if this.Next.Next == nil {
		this.Next = nil		
		return
	}
}

type MyLinkedList struct {
	Node *MyNode
	Size int
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	n := new(MyNode)
	n.Next = n
	n.Prev = n
	l := MyLinkedList{
		Node: n,
		Size: 0,
	}
	return l
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	if index >= this.Size || index < 0 {
		return -1
	}
	iter := this.Node
	for i := 0; i <= index && iter != nil; i, iter = i+1, iter.Next {
		if i == index {
			return iter.Val
		}
	}
	return -1
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {

	n := new(MyNode)
	n.Val = val
	if this.Size == 0 {
		this.Node = n
	} else {
		this.Node.Prev = n
		n.Next = this.Node
		this.Node = n
	}
	this.Size++
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	iter := this.Node
	for ; iter.Next != nil; iter = iter.Next {

	}
	iter.Push(val)
	this.Size++
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {

	if this.Size < index || index < 0 {
		return
	}
	n := new(MyNode)
	n.Val = val

	if this.Size == 0 {
		this.Node = n
	} else {
		iter := this.Node
		for i := 0; i < index; i, iter = i+1, iter.Next {
			if i == index-1 {
				iter.Push(val)
			}
		}
	}
	this.Size++

}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if this.Size < index || index < 0 {
		return
	}
	if index == 0 {
		this.Node = this.Node.Next

	} else {
		iter := this.Node
		for i := 0; i < index;  i++{
			iter = iter.Next
		}
		iter.PopNext()
	}
	this.Size--
}

func (t *MyLinkedList) PrintALL() {
	iter := t.Node
	for ; iter != nil; iter = iter.Next {
		print(iter.Val)
		print(" ,")
	}
	print("Size: ")
	print(t.Size)
	println()
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
