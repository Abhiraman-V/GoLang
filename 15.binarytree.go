package main

import (
	"fmt"
)

/*
	Binary Tree
		A tree is a representation of a hierarchical structure. a non-sequential data structure.
		A binary tree is a tree where every node has max 2 children.
		A binary search tree has the property of the left node having a value less than the value on the right node.

		Root: the level 0 of the tree
		Child: each node of the tree that’s not the root
		Internal node: each node with at least a child
		Leaf: each node that has no children
		Subtree: the set of node with a certain node as root
*/

type Tree struct {
	root *Node
}

type Node struct {
	key   byte
	data  int
	left  *Node
	right *Node
}

func (t *Tree) insert(data byte) {
	if t.root == nil {
		t.root = &Node{key: data}
	} else {
		t.root.insert(data)
	}
}

func (n *Node) insert(data byte) {
	if data <= n.key {
		if n.left == nil {
			n.left = &Node{key: data}
		} else {
			n.left.insert(data)
		}
	} else {
		if n.right == nil {
			n.right = &Node{key: data}
		} else {
			n.right.insert(data)
		}
	}
}

func printPreOrder(n *Node) {
	if n == nil {
		return
	} else {
		fmt.Printf("%d ", n.data)
		printPreOrder(n.left)
		printPreOrder(n.right)
	}
}

func printPostOrder(n *Node) {
	if n == nil {
		return
	} else {
		printPostOrder(n.left)
		printPostOrder(n.right)
		fmt.Printf("%c ", n.data)
	}
}

func printInOrder(n *Node) {
	if n == nil {
		return
	} else {
		printInOrder(n.left)
		fmt.Printf("%c ", n.key)
		printInOrder(n.right)
	}
}

func main() {
	var t Tree

	t.insert('K')
	t.insert('L')
	t.insert('O')
	t.insert('O')
	t.insert('D')
	t.insert('O')
	t.insert('N')
	t.insert('E')
	t.insert('H')

	fmt.Printf("Pre Order: ")
	printPreOrder(t.root)
	fmt.Println()
	fmt.Printf("Post Order: ")
	printPostOrder(t.root)
	fmt.Println()
	fmt.Printf("In Order: ")
	printInOrder(t.root)
	fmt.Println()
}
