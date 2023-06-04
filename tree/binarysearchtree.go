package tree

import (
	"fmt"

	"github.com/4d3v/golangdatastructure/utils"
)

const (
	INT    = iota // 0
	STRING        // 1
)

type (
	BinarySearchTree struct {
		root       *node
		comparator utils.Comparator
		length     int
		keyType    interface{}
	}

	node struct {
		key   interface{}
		value interface{}
		left  *node
		right *node
	}
)

func IntBST() *BinarySearchTree {
	return new(utils.IntComparator, INT)
}

func StringBST() *BinarySearchTree {
	return new(utils.StringComparator, STRING)
}

func new(comparator utils.Comparator, keyType interface{}) *BinarySearchTree {
	return &BinarySearchTree{nil, comparator, 0, keyType}
}

func (t *BinarySearchTree) Len() int {
	return t.length
}

func (t *BinarySearchTree) Insert(key, val interface{}) {
	n := &node{t.assertKeyType(key), val, nil, nil}
	if t.length == 0 {
		t.root = n
	} else {
		insertNode(n, t.root, t.comparator)
	}
	t.length++
}

func (t *BinarySearchTree) search(key interface{}) *node {
	cur := t.root
	for cur != nil {
		if t.comparator(key, cur.key) == 0 {
			return cur
		}
		if t.comparator(key, cur.key) == -1 {
			cur = cur.left
		} else if t.comparator(key, cur.key) == 1 {
			cur = cur.right
		}
	}
	return nil
}

func (t *BinarySearchTree) GetNode(value interface{}) *node {
	return t.search(value)
}

func (t *BinarySearchTree) GetNodeValue(key interface{}) (value interface{}, found bool) {
	n := t.search(key)
	if n != nil {
		return n.value, true
	}
	return nil, false
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
		fmt.Println(node.key)

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

func (t *BinarySearchTree) assertKeyType(key interface{}) interface{} {
	switch t.keyType {
	case INT:
		k, ok := key.(int)
		if !ok {
			panic("Type error: expected an int")
		}
		return k
	case STRING:
		k, ok := key.(string)
		if !ok {
			panic("Type error: expected a string")
		}
		return k
	default:
		panic("Type error: expected an integer or string")
	}
}

func insertNode(newNode, curNode *node, comparator utils.Comparator) {
	if comparator(newNode.key, curNode.key) == 0 {
		return
	}

	if comparator(newNode.key, curNode.key) == -1 {
		if curNode.left == nil {
			curNode.left = newNode
			return
		}
		insertNode(newNode, curNode.left, comparator)
	} else if comparator(newNode.key, curNode.key) == 1 {
		if curNode.right == nil {
			curNode.right = newNode
			return
		}
		insertNode(newNode, curNode.right, comparator)
	}
}
