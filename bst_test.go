package gocol

import (
	"testing"
)

func TestBSTInsert(t *testing.T) {
	bst := NewBST[int, string]()
	bst.Insert(5, "a")
	bst.Insert(7, "b")
	bst.Insert(3, "c")
	bst.Insert(6, "d")
	bst.Insert(2, "e")

	if min := bst.Min(); min.key != 2 {
		t.Fail()
	}
	if max := bst.Max(); max.key != 7 {
		t.Fail()
	}
	if tr := bst.Search(6); tr.val != "d" {
		t.Fail()
	}
}
