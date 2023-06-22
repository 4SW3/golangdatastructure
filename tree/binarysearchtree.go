package tree

import (
	"fmt"

	"github.com/4d3v/golangdatastructure/queue"
	"github.com/4d3v/golangdatastructure/utils"
)

const (
	INT    = iota // 0
	STRING        // 1
)

const (
	PreOrder TraversalType = iota
	InOrder
	PostOrder
)

type TraversalType int

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

func (t *BinarySearchTree) hasOnlyRoot(n1, n2 *node) bool {
	return t.length == 1 && n1 == n2
}

func (t *BinarySearchTree) isRoot(n *node) bool {
	return n == t.root
}

func (t *BinarySearchTree) isLeaf(n *node) bool {
	return n.left == nil && n.right == nil
}

func (t *BinarySearchTree) hasOneChild(n *node) bool {
	return (n.left != nil && n.right == nil) || (n.left == nil && n.right != nil)
}

func (t *BinarySearchTree) hasTwoChildren(n *node) bool {
	return n.left != nil && n.right != nil
}

func (t *BinarySearchTree) deleteNode(key interface{}) *node {
	node, parent := t.GetNodeAndParent(key)
	if node == nil {
		return nil
	}

	if t.hasOnlyRoot(node, parent) {
		t.root = nil
		return node
	}

	if t.isLeaf(node) {
		if parent.left == node {
			parent.left = nil
		} else {
			parent.right = nil
		}
		return node
	}

	if t.hasTwoChildren(node) {
		min, minP := t.findMinWithParent(node.right)

		// replace node with min
		if node.left != min {
			min.left = node.left
		}
		if node.right != min {
			min.right = node.right
		}

		// remove pointer from minP to min
		if minP.left == min {
			minP.left = nil
		} else {
			minP.right = nil
		}

		// replace parent pointer to node with min
		if t.isRoot(node) {
			t.root = min
		} else {
			if parent.left == node {
				parent.left = min
			} else {
				parent.right = min
			}
		}

		return node
	}

	if t.hasOneChild(node) {
		if node.left != nil {
			if parent.left == node {
				parent.left = node.left
			} else {
				parent.right = node.left
			}
		} else {
			if parent.left == node {
				parent.left = node.right
			} else {
				parent.right = node.right
			}
		}
		return node
	}

	return nil
}

func (t *BinarySearchTree) Delete(key interface{}) *node {
	if n := t.deleteNode(key); n != nil {
		t.length--
		return &node{n.key, n.value, nil, nil}
	}
	return nil
}

func (t *BinarySearchTree) BFS() []interface{} {
	var data []interface{}
	q := queue.New()
	q.Enqueue(t.root)

	for q.Len() > 0 {
		n := q.Dequeue().(*node)
		data = append(data, n.key)

		if n.left != nil {
			q.Enqueue(n.left)
		}
		if n.right != nil {
			q.Enqueue(n.right)
		}
	}

	return data
}

func (t *BinarySearchTree) DFSPreOrder() []interface{} {
	var visited []interface{}
	n := t.root

	// var traverse func(n *node)
	// traverse = func(n *node) {
	// 	visited = append(visited, n.key)
	// 	if n.left != nil {
	// 		traverse(n.left)
	// 	}
	// 	if n.right != nil {
	// 		traverse(n.right)
	// 	}
	// }
	// traverse(n)
	traverse(n, PreOrder, &visited)

	return visited
}

func (t *BinarySearchTree) DFSInOrder() []interface{} {
	var visited []interface{}
	n := t.root

	// var traverse func(n *node)
	// traverse = func(n *node) {
	// 	if n.left != nil {
	// 		traverse(n.left)
	// 	}
	// 	visited = append(visited, n.key)
	// 	if n.right != nil {
	// 		traverse(n.right)
	// 	}
	// }
	// traverse(r)
	traverse(n, InOrder, &visited)

	return visited
}

func (t *BinarySearchTree) DFSPostOrder() []interface{} {
	var visited []interface{}
	n := t.root

	// var traverse func(n *node)
	// traverse = func(n *node) {
	// 	if n.left != nil {
	// 		traverse(n.left)
	// 	}
	// 	if n.right != nil {
	// 		traverse(n.right)
	// 	}
	// 	visited = append(visited, n.key)
	// }
	// traverse(r)
	traverse(n, PostOrder, &visited)

	return visited
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

func traverse(n *node, t TraversalType, visited *[]interface{}) {
	if n == nil {
		return
	}

	switch t {
	case PreOrder:
		*visited = append(*visited, n.key)
		traverse(n.left, t, visited)
		traverse(n.right, t, visited)

	case InOrder:
		traverse(n.left, t, visited)
		*visited = append(*visited, n.key)
		traverse(n.right, t, visited)

	case PostOrder:
		traverse(n.left, t, visited)
		traverse(n.right, t, visited)
		*visited = append(*visited, n.key)
	}
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
