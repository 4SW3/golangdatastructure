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

func (t *BinarySearchTree) GetNodeAndParent(key interface{}) (node, parent *node) {
	if t.length == 0 {
		return nil, nil
	}
	if t.length == 1 || t.comparator(key, t.root.key) == 0 {
		return t.root, t.root
	}
	return t.searchWithParent(key)
}

func isLeaf(n *node) bool {
	return n.left == nil && n.right == nil
}

// TODO
// fix bug when removing parent
// change func to return *node
func (t *BinarySearchTree) Delete(key interface{}) {
	n, p := t.GetNodeAndParent(key)

	if n == nil {
		return
	}

	if t.Len() == 1 && p == nil {
		t.root = nil
		return
	}

	if isLeaf(n) {
		if p.left == n {
			p.left = nil
		} else {
			p.right = nil
		}
		return
	}

	if n.left != nil && n.right != nil {
		min, minParent := t.findMinWithParent(n.right)
		fmt.Println("A")

		if minParent.left == min {
			minParent.left = nil
		} else {
			minParent.right = nil
		}

		if p.left == n {
			p.left = min
		} else {
			p.right = min
		}

		if n.left != min {
			min.left = n.left
		}
		if n.right != min {
			min.right = n.right
		}

		n.left = nil
		n.right = nil
		fmt.Println("B")

		return
	}

	// left or right has data
	if n.left != nil {
		if p.left == n {
			p.left = n.left
		} else {
			p.right = n.left
		}
	} else {
		if p.left == n {
			p.left = n.right
		} else {
			p.right = n.right
		}
	}
}

func (t *BinarySearchTree) PrintBST() {
	printBT("", t.root, false)
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

func (t *BinarySearchTree) searchWithParent(key interface{}) (node, parent *node) {
	cur := t.root
	par := cur
	for cur != nil {
		if t.comparator(key, cur.key) == 0 {
			return cur, par
		}
		if t.comparator(key, cur.key) == -1 {
			par = cur
			cur = cur.left
		} else if t.comparator(key, cur.key) == 1 {
			par = cur
			cur = cur.right
		}
	}
	return nil, nil
}

// find min with parent
func (t *BinarySearchTree) findMinWithParent(n *node) (node, parent *node) {
	cur := n
	par := cur
	for cur.left != nil {
		par = cur
		cur = cur.left
	}
	return cur, par
}

// find max with parent
func (t *BinarySearchTree) findMaxWithParent(n *node) (node, parent *node) {
	cur := n
	par := cur
	for cur.right != nil {
		par = cur
		cur = cur.right
	}
	return cur, par
}

// find min node
func (t *BinarySearchTree) FindMin() *node {
	cur := t.root
	for cur.left != nil {
		cur = cur.left
	}
	return cur
}

// find max node
func (t *BinarySearchTree) FindMax() *node {
	cur := t.root
	for cur.right != nil {
		cur = cur.right
	}
	return cur
}

// func (t *BinarySearchTree) invertHelper(n *node) *node {
// 	if n == nil {
// 		return nil
// 	}
// 	if n.left != nil {
// 		t.invertHelper(n.left)
// 	}
// 	if n.right != nil {
// 		t.invertHelper(n.right)
// 	}

// 	n.left, n.right = n.right, n.left
// 	return n
// }
