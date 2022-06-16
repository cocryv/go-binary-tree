package main

import "fmt"

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

type Tree struct {
	Root *Node
	len  int
}

func (t *Tree) Add(i int) {
	t.AddByNode(i, t.Root)
	t.len++
}

func (t *Tree) AddByNode(i int, n *Node) {
	if i < n.Key {
		if n.Left == nil {
			n.Left = &Node{Key: i}
		} else {
			t.AddByNode(i, n.Left)
		}

	} else {
		if n.Right == nil {
			n.Right = &Node{Key: i}
		} else {
			t.AddByNode(i, n.Right)
		}

	}
}

func (t *Tree) Search(i int) bool {
	return t.SearchByNode(i, t.Root)
}

func (t *Tree) SearchByNode(i int, n *Node) bool {

	if n == nil {
		return false
	}

	if n.Key < i {
		//move right
		return t.SearchByNode(i, n.Right)

	} else if n.Key > i {
		// move left
		return t.SearchByNode(i, n.Right)
	}
	return true
}

func main() {

	tree := Tree{Root: &Node{Key: 4}}
	tree.Add(2)
	tree.Add(1)
	tree.Add(6)
	tree.Add(5)
	tree.Add(7)
	fmt.Println(tree.Root)
	fmt.Println(tree.Root.Right.Right.Key)
}
