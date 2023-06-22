package main

import (
	"fmt"

	"github.com/4d3v/golangdatastructure/tree"
)

type desc struct {
	id   int
	desc string
}

func main() {
	// t := tree.IntBST()
	// t.Insert(8, desc{1, "this is number 8"})
	// t.Insert(10, desc{2, "this is number 10"})
	// t.Insert(5, desc{3, "this is number 5"})
	// t.Insert(11, desc{4, "this is number 11"})
	// t.Insert(2, desc{5, "this is number 2"})
	// t.Insert(9, desc{6, "this is number 9"})
	// t.Insert(6, desc{7, "this is number 6"})
	// t.PrintBST()

	// bst := tree.StringBST()
	// bst.Insert("orange", desc{1, "this is an orange"})
	// bst.Insert("banana", desc{2, "this is a banana"})
	// bst.Insert("apple", desc{3, "this is an apple"})
	// bst.Insert("cherry", desc{4, "this is a cherry"})
	// bst.Insert("cherryx", desc{41, "this is a cherry"})
	// bst.Insert("cherra", desc{42, "this is a cherry"})
	// bst.Insert("pineapple", desc{5, "this is a pineapple"})
	// bst.Insert("peach", desc{6, "this is a peach"})
	// bst.Insert("watermelon", desc{7, "this is a watermelon"})

	bst := tree.IntBST()
	bst.Insert(15, desc{1, "Number"})
	bst.Insert(1, desc{2, "Number"})
	bst.Insert(76, desc{3, "Number"})
	bst.Insert(7, desc{4, "Number"})
	bst.Insert(56, desc{5, "Number"})
	bst.Insert(80, desc{6, "Number"})
	bst.Insert(19, desc{7, "Number"})
	bst.Insert(65, desc{8, "Number"})
	fmt.Println(bst.DFSPostOrder())
}

// └──8
//    ├──5
//    │   ├──2
//    │   └──6
//    └──10
//        ├──9
//        └──11

// └──orange
//     ├──banana
//     │   ├──apple
//     │   └──cherry
//     └──pineapple
//         ├──peach
//         └──watermelon
