package main

import "fmt"

type BinaryTree struct {
	root *Node
}

type Node struct {
	key   int
	left  *Node
	right *Node
}


func(b *BinaryTree) insert(key int) {
	if b.root == nil {
		b.root = &Node{key: key}
	} else {
		b.root.insert(key)
	}
}

func(n *Node) insert(key int) {
	if n.key > key { // insert Left
		if n.left == nil {
			n.left = &Node{key: key}

		} else {
			n.left.insert(key)
		}
	} else if n.key < key { // insert Right
		if n.right ==  nil {
			n.right = &Node{key: key}
		} else {
			n.right.insert(key)
		}
	}
}

func main() {
	b := BinaryTree{}
	c:= []int{9,6,8,12,15,2,7,14,20,11,1}
	for i:=0; i<len(c); i++ {
		b.insert(c[i])
	}
	b.InOrder()
	fmt.Println()
	b.PreOrder()
	fmt.Println()
	b.PostOrder()
}

func(b *BinaryTree) InOrder()  { // left root right
	if b.root == nil {
		return
	} else {
		b.root.inOrder()
	}
}

func(b *BinaryTree) PreOrder()  { // root left right
	if b.root == nil {
		return
	} else {
		b.root.preOrder()
	}
}

func(b *BinaryTree) PostOrder()  { // left right root
	if b.root == nil {
		return
	} else {
		b.root.postOrder()
	}
}

func(n *Node) inOrder()  { // left root right
	if n.left != nil {
		n.left.inOrder()
	}
	fmt.Print(n.key," ")
	if n.right != nil{
		n.right.inOrder()
	}
}

func(n *Node) preOrder() { // root left right
	if n != nil {
		fmt.Print(n.key," ")
		if n.left != nil {
			n.left.preOrder()
		}
		if n.right != nil {
			n.right.preOrder()
		}
	}
}

func (n *Node) postOrder() { // left right root
	if n.left != nil {
		n.left.postOrder()
	}
	if n.right != nil {
		n.right.postOrder()
	}
	fmt.Print(n.key," ")
}


func (b *BinaryTree) search(key int) bool {
	 if b.root == nil {
	 	return false
	 } else {
	 	return b.root.search(key)
	 }
	 return true
}

func(n *Node) search(key int) bool {
	if n.key > key { // search left
		if n.left == nil {
			return false
		} else {
			return n.left.search(key)
		}
	} else if n.key < key { // search right
		if n.right == nil {
			return false
		} else {
			return n.right.search(key)
		}
	}
	return true
}

func(b *BinaryTree) Delete(key int) *Node {
	if b.root == nil {
		return nil
	} else {
		return b.root.Delete(key)
	}
}

func(n *Node) Delete(key int) *Node {
	if n.key < key {
		n.right = n.right.Delete(key)
	} else if n.key > key {
		n.left = n.left.Delete(key)
	} else {
		if n.left == nil {
			return n.right
		} else if n.right == nil {
			return n.left
		}
		min := n.right.Min()
		n.key = min
		n.right = n.right.Delete(min)
	}
	return n
}


func(n *Node) Min() int {
	if n.left == nil {
		return n.key
	}
	return n.left.Min()
}