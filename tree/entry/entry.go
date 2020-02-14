package main

import (
	"fmt"
	"function/tree"
)

type myNode struct {
	node *tree.Node
}

func (my *myNode) postTraverse() {
	if my == nil || my.node == nil {
		return
	}

	left := myNode{my.node.Left}
	right := myNode{my.node.Right}
	left.postTraverse()
	right.postTraverse()
	my.node.Print()
}

func main() {
	var root tree.Node
	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)
	root.Right.Left.SetValue(4)
	root.Traverse()

	myroot := myNode{&root}
	myroot.postTraverse()

	nodeCount := 0
	root.TraverseFunc(func(node *tree.Node) {
		nodeCount++
	})
	fmt.Println("nodeCount:", nodeCount)

	fmt.Println()
	maxNode := 0
	root.TraverseFunc(func(node *tree.Node) {
		if node.Value > maxNode {
			maxNode = node.Value
		}
	})
	fmt.Println("maxNodeValue with functional :", maxNode)

	fmt.Println()

	c := root.TraverseWithChannel()
	maxNodeChan := 0
	for node := range c {
		if node.Value > maxNodeChan {
			maxNodeChan = node.Value
		}
	}
	fmt.Println("maxNodeValue with channel :", maxNodeChan)

}
