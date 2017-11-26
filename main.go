package main

import "fmt"

type Edge struct {
	start *Node
	end *Node
}

type Node struct {
	id string
}

func createNode(id string) *Node {
	node := new(Node)
	node.id = id
	return node
}

func createEdge(start *Node, end *Node) *Edge {
	edge := new(Edge)
	edge.start = start
	edge.end = end
	return edge
}

func isCyclic(edges []*Edge) bool {
	fmt.Printf("%s %s\n", edges[2].start.id, edges[2].end.id)

	for i:= 0; i < len(edges); i++ {
		var start= edges[i].start
		if startJourney(edges, start, edges[i], i) {
			return true
		}
	}

	return false
}

func startJourney(edges []*Edge, start *Node, current *Edge, pos int) bool {
	if start == current.end {
		return true
	}

	pos++

	if pos >= len(edges) {
		return false
	}
	return startJourney(edges, start, edges[pos], pos)
}

func main() {

	nodeA := createNode("deniss")
	nodeB := createNode("daven")
	nodeC := createNode("antons")
	nodeD := createNode("serge")
	nodeE := createNode("jevk")

	fmt.Printf("%s %s %s %s %s \n", nodeA.id, nodeB.id, nodeC.id, nodeD.id, nodeE.id)

	var edges [5]*Edge

	edges[0] = createEdge(nodeA, nodeB)
	edges[1] = createEdge(nodeB, nodeC)
	edges[2] = createEdge(nodeC, nodeE)
	edges[3] = createEdge(nodeA, nodeE)
	edges[4] = createEdge(nodeB, nodeE)

	fmt.Printf("The result is: %t \n", isCyclic(edges[:]))
}

