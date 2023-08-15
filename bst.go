package gocol

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type BST[K constraints.Ordered, V any] struct {
	root  *BST[K, V]
	left  *BST[K, V]
	right *BST[K, V]
	key   K
	val   V
}

func NewBST[K constraints.Ordered, V any]() *BST[K, V] {
	var key K
	var val V
	return &BST[K, V]{
		root:  nil,
		left:  nil,
		right: nil,
		key:   key,
		val:   val,
	}
}

func (bst *BST[K, V]) Insert(key K, val V) {
	var travPtr *BST[K, V] = bst.root
	var rootPtr *BST[K, V] = nil
	leaf := &BST[K, V]{
		root:  nil,
		left:  nil,
		right: nil,
		key:   key,
		val:   val,
	}

	// Traverse tree to correct leaf
	for travPtr != nil {
		rootPtr = travPtr
		if leaf.key < travPtr.key {
			travPtr = travPtr.left
		} else {
			travPtr = travPtr.right
		}
	}
	leaf.root = rootPtr

	// Attach new leaf node
	if rootPtr == nil {
		// Empty tree case
		bst.root = leaf
	} else if leaf.key < rootPtr.key {
		rootPtr.left = leaf
	} else {
		rootPtr.right = leaf
	}
}

func (bst *BST[K, V]) Min() *BST[K, V] {
	travPtr := bst.root
	for travPtr.left != nil {
		travPtr = travPtr.left
	}
	return travPtr
}

func (bst *BST[K, V]) Max() *BST[K, V] {
	travPtr := bst.root
	for travPtr.right != nil {
		travPtr = travPtr.right
	}
	return travPtr
}

func (bst *BST[K, V]) Search(key K) *BST[K, V] {
	travPtr := bst.root
	for travPtr != nil && travPtr.key != key {
		if key < travPtr.key {
			travPtr = travPtr.left
		} else {
			travPtr = travPtr.right
		}
	}
	return travPtr
}

func (bst *BST[K, V]) Print() {
	inorderWalk[K, V](bst.root)
}

func inorderWalk[K constraints.Ordered, V any](t *BST[K, V]) {
	if t != nil {
		inorderWalk[K, V](t.left)
		fmt.Print(t.key, " ")
		inorderWalk[K, V](t.right)
	}
}
