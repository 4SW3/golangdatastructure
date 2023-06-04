package tree

import (
	"fmt"

	"github.com/4d3v/golangdatastructure/utils"
)

type (
	BinarySearchTree struct {
		root   *node
		length int
	}

	node struct {
		value interface{}
		left  *node
		right *node
	}
)

func New() *BinarySearchTree {
	return &BinarySearchTree{nil, 0}
}

func (t *BinarySearchTree) Len() int {
	return t.length
}

func (t *BinarySearchTree) Insert(val interface{}) {
	n := &node{val, nil, nil}

	if t.length == 0 {
		t.root = n
		t.length++
		return
	}

	for cur := t.root; cur != nil; {

		if utils.IntComparator(val, cur.value) == 0 {
			return
		}

		if utils.IntComparator(val, cur.value) == -1 {
			if cur.left == nil {
				cur.left = n
				t.length++
				return
			}
			cur = cur.left
		} else if utils.IntComparator(val, cur.value) == 1 {
			if cur.right == nil {
				cur.right = n
				t.length++
				return
			}
			cur = cur.right
		}
	}
}

func printBT(prefix string, node *node, isLeft bool) {
	if node != nil {
		fmt.Print(prefix)

		if isLeft {
			fmt.Print("├──")
		} else {
			fmt.Print("└──")
		}

		// print the value of the node
		fmt.Println(node.value)

		// enter the next tree level - left and right branch
		printBT(prefix+getBranchPrefix(isLeft), node.left, true)
		printBT(prefix+getBranchPrefix(isLeft), node.right, false)
	}
}

func getBranchPrefix(isLeft bool) string {
	if isLeft {
		return "│   "
	}
	return "    "
}

func (t *BinarySearchTree) PrintBST() {
	printBT("", t.root, false)
}
