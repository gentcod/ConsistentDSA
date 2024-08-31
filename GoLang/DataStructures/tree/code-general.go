package tree

import "fmt"

type Data struct {
	Data   any
	Weight int
}

type Node[T Data] struct {
	Data   *Data
	parent *Node[T] // for reference when traversing the tree
	LChild *Node[T]
	RChild *Node[T]
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
		if n.LChild == nil {
			n.LChild = &Node[T]{Data: &value}
			if n.parent != nil && n.parent.LChild == nil {
				n.LChild.parent = n.parent
			} else {
				n.LChild.parent = n
			} 
			// fmt.Println("left-child: ", n.LChild.Data, "parent: ", n.LChild.parent.Data)
		} else {
			n.LChild.Insert(value)
		}
	}

	if value.Weight > n.Data.Weight {
		if n.RChild == nil {
			n.RChild = &Node[T]{Data: &value}
			if n.parent != nil && n.parent.RChild == nil {
				n.RChild.parent = n.parent
			} else {
				n.RChild.parent = n
			}
			// fmt.Println("right-child: ", n.RChild.Data, "parent: ", n.RChild.parent.Data)
		} else {
			n.RChild.Insert(value)
		}
	}
}

// Diplay traverses the tree and shows the data in the LinkedList
func (n *Node[T]) Diplay(pos string) {
	fmt.Printf("Weight: %v, Position: %v, Data: %v\n", n.Data.Weight, pos, n.Data.Data)

	if n.LChild != nil {
		newPos := "left; Parent: %v"
		n.LChild.Diplay(fmt.Sprintf(newPos, n.LChild.parent.Data.Weight))
	}

	if n.RChild != nil{
		newPos := "right; Parent: %v"
		n.RChild.Diplay(fmt.Sprintf(newPos, n.RChild.parent.Data.Weight))
	}
}
