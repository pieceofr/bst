package bst

import (
	//"errors"
	//"log"
	//"sync"
	"fmt"
)

/*implemt a binary search tree with key type integer*/
type Node struct {
	key   int
	left  *Node
	right *Node
}

type BST struct {
	root *Node
}

func (b *BST) Insert(key int) error {
	node := Node{key: key}
	b.root = insertNode(&node, b.root)
	return nil
}

func insertNode(new *Node, p *Node) *Node {
	/* ending condition: p is where new node should be*/
	if p == nil {
		return new
	}
	/*new node is duplicated p = p*/
	if new.key == p.key {
		return p
	}

	if new.key < p.key {
		p.left = insertNode(new, p.left)
	} else {
		p.right = insertNode(new, p.right)
	}
	return p
}

func (b *BST) PreOrderPrint() []int {
	preList := make([]int, 0, 1000)
	preList = preOrderHelper(b.root, preList)
	return preList
}

func preOrderHelper(r *Node, l []int) []int {
	if r != nil {
		fmt.Printf("%d ", r.key)
		l = append(l, r.key)
		// because the stack of recursive call 
		// stores the old slice so we need to make the l the last append list
		l = preOrderHelper(r.left, l)
		l = preOrderHelper(r.right, l)
	}
	return l
}
