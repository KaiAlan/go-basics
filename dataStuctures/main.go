package main

import (
	"github/KaiAlan/basics/dataStuctures/Bst"
)

func main() {
	bst := Bst.BSTWrapper{}

	bst.Insert(8)
	bst.Insert(4)
	bst.Insert(10)
	bst.Insert(1)
	bst.Insert(6)
	bst.Insert(2)
	bst.Insert(5)

	bst.InOrderTraversal()
	bst.PreOrderTraversal()
	bst.PostOrderTraversal()

	bst.String()

	bst.Search(2)

	bst.Remove(4)

	bst.String()

	bst.InOrderTraversal()
	bst.PreOrderTraversal()
	bst.PostOrderTraversal()
}
