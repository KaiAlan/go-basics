package Bst

import (
	"fmt"
)

type TreeNode struct {
	value int
	left  *TreeNode
	right *TreeNode
}

type BSTWrapper struct {
	root   *TreeNode
	length int
}

func (bst *BSTWrapper) Insert(value int) {
	node := &TreeNode{value, nil, nil}

	if bst.root == nil {
		bst.root = node
	} else {
		insertNode(bst.root, node)
	}
	bst.length++
}

func insertNode(node, newNode *TreeNode) {
	if newNode.value < node.value {
		if node.left == nil {
			node.left = newNode
		} else {
			insertNode(node.left, newNode)
		}
	} else {
		if node.right == nil {
			node.right = newNode
		} else {
			insertNode(node.right, newNode)
		}
	}
}

func (bst *BSTWrapper) InOrderTraversal() {
	if bst.root != nil {
		fmt.Printf("In Order Traversal:\t")
		inOrderTraverse(bst.root)
		fmt.Printf("\n")
	}
}

func inOrderTraverse(node *TreeNode) {
	if node != nil {
		inOrderTraverse(node.left)
		fmt.Printf("%d ", node.value)
		inOrderTraverse(node.right)
	}
}

func (bst *BSTWrapper) PreOrderTraversal() {
	if bst.root != nil {
		fmt.Printf("Pre Order Traversal:\t")
		preOrderTraverse(bst.root)
		fmt.Printf("\n")
	}
}

func preOrderTraverse(node *TreeNode) {
	if node != nil {
		fmt.Printf("%d ", node.value)
		preOrderTraverse(node.left)
		preOrderTraverse(node.right)
	}
}

func (bst *BSTWrapper) PostOrderTraversal() {
	if bst.root != nil {
		fmt.Printf("Post Order Traversal:\t")
		postOrderTraverse(bst.root)
		fmt.Printf("\n")
	}
}

func postOrderTraverse(node *TreeNode) {
	if node != nil {
		postOrderTraverse(node.left)
		postOrderTraverse(node.right)
		fmt.Printf("%d ", node.value)
	}
}

func (bst *BSTWrapper) Search(value int) {

	fmt.Printf("Searching for value: %d\t | ", value)
	if search(bst.root, value) {
		fmt.Printf("Found it.\n")
	} else {
		fmt.Printf("Value %d is not in the Tree.\n", value)
	}
}

func search(n *TreeNode, key int) bool {
	if n == nil {
		return false
	}
	if key < n.value {
		return search(n.left, key)
	}
	if key > n.value {
		return search(n.right, key)
	}
	return true
}

func (bst *BSTWrapper) Remove(value int) {
	fmt.Printf("Attempting to remove value: %d\t", value)

	if remove(bst.root, value) != nil {
		fmt.Printf("Value %d removed successfully!\n", value)
	} else {
		fmt.Printf("Value %d not found!\n", value)
	}
}

func remove(node *TreeNode, value int) *TreeNode {
	if node == nil {
		return nil
	}
	if value < node.value {
		return remove(node.left, value)
	}
	if value > node.value {
		return remove(node.right, value)
	}

	if node.left == nil && node.right == nil {
		node = nil
		return nil
	}

	if node.left == nil {
		node = node.right
		return node
	}

	if node.right == nil {
		node = node.left
		return node
	}

	lastmostRightnode := node.right

	for {
		if lastmostRightnode != nil && lastmostRightnode.left != nil {
			lastmostRightnode = lastmostRightnode.left
		} else {
			break
		}
	}

	node.value = lastmostRightnode.value
	node.right = remove(node.right, node.value)

	return node
}

func (bst *BSTWrapper) String() {
	fmt.Println("------------------------------------------------")
	stringify(bst.root, 0)
	fmt.Println("------------------------------------------------")
}

func stringify(node *TreeNode, level int) {
	if node != nil {
		format := ""
		for i := 0; i < level; i++ {
			format += "       "
		}
		format += "---[ "
		level++
		stringify(node.right, level)
		fmt.Printf(format+"%d\n", node.value)
		stringify(node.left, level)
	}
}
