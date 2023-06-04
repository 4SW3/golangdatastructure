package main

import (
	"github.com/4d3v/golangdatastructure/tree"
)

func main() {
	t := tree.New()
	t.Insert(8)
	t.Insert(10)
	t.Insert(5)
	t.Insert(11)
	t.Insert(2)
	t.Insert(9)
	t.Insert(6)
	t.Insert(7)
	t.Insert(1)

	t.PrintBST()

}

// └──8
//    ├──5
//    │   ├──2
//    │   └──6
//    └──10
//        ├──9
//        └──11
