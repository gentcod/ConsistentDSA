package tree

import "fmt"

// type BSTTree[T comparable] struct {
// 	Root     *T
// 	count    int
// 	vertices int
// }

type Data struct {
	Data   any
	Weight int
}

type Node[T Data] struct {
	Data   *Data
	parent *Node[T] // for reference when traversing the tree
	Left *Node[T]
	Right *Node[T]
}

type BSTTree[T Data] struct {
	Root     *Node[T]
	count    int
	vertices int
}

func New[T Data](value Data) *BSTTree[T] {
	tree := BSTTree[T]{}
	node := &Node[T]{Data: &value}
	tree.Root = node
	tree.count++
	return &tree
}

func (t *BSTTree[T]) Add(values ...Data) {
	for _, value := range values {
		t.Root.Insert(value)
		t.count++

		if t.count%2 == 1 {
			t.vertices++
		}
	}
}

func (n *Node[T]) Insert(value Data) {
	if n == nil {
		return
	}

	if value.Weight < n.Data.Weight {
		if n.Left == nil {
			n.Left = &Node[T]{Data: &value}
			if n.parent != nil && n.parent.Left == nil {
				n.Left.parent = n.parent
			} else {
				n.Left.parent = n
			} 
		} else {
			n.Left.Insert(value)
		}
	}

	if value.Weight > n.Data.Weight {
		if n.Right == nil {
			n.Right = &Node[T]{Data: &value}
			if n.parent != nil && n.parent.Right == nil {
				n.Right.parent = n.parent
			} else {
				n.Right.parent = n
			}
		} else {
			n.Right.Insert(value)
		}
	}
}

// Diplay traverses the tree and shows the data in the LinkedList
func (n *Node[T]) Diplay(pos string) {
	fmt.Printf("Weight: %v, Position: %v, Data: %v\n", n.Data.Weight, pos, n.Data.Data)

	if n.Left != nil {
		newPos := "left; Parent: %v"
		n.Left.Diplay(fmt.Sprintf(newPos, n.Left.parent.Data.Weight))
	} 
	
	if n.Right != nil {
		newPos := "right; Parent: %v"
		n.Right.Diplay(fmt.Sprintf(newPos, n.Right.parent.Data.Weight))
	}
}

// SAMPLE DATA
// data := []tree.Data{ 
// 	{Data: "hi", Weight: 5}, 
// 	{Data: "hey", Weight: 16}, 
// 	{Data: "tudo", Weight: 18},
// 	{Data: "try", Weight: 10},
// 	{Data: "holla", Weight: 8},
// 	{Data: "holla", Weight: 17},
// }
