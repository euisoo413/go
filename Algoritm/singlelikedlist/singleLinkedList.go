package singlelinkedlist

type Node[T any] struct {
	next  *Node[T]
	Value T
}

type LinkedList[T any] struct {
	root *Node[T]
	tail *Node[T]

	count int
}

func (l *LinkedList[T]) PushBack(value T) {
	node := &Node[T]{
		Value: value,
	}
	l.count++
	if l.root == nil {
		l.root = node
		l.tail = node
		return
	}

	l.tail.next = node
	l.tail = node
}

func (l *LinkedList[T]) PushFront(value T) {
	node := &Node[T]{
		Value: value,
	}
	l.count++
	if l.root == nil {
		l.root = node
		l.tail = node
		return
	}

	node.next = l.root
	l.root = node
}

// package singlelinkedlist

// import "fmt"

// type Node[T any] struct{
// 	next *Node[T]
// 	Value T
// }

// type LinkedList[T any] struct{
// 	root *Node[T]
// 	tail *Node[T]
// }

// func (l *LinkedList[T]) PushBack(value T){
// 	node := &Node[T]{
// 		Value: value,
// 	}
// 	if l.root==nil{
// 		l.root=node
// 		l.tail=node
// 		return
// 	}


// }


// func main(){


// }