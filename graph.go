package shortpath

import (
	"strconv"
)

type Graph struct {
	Nodes []*node
}

type  node struct {
	name string
	graph *Graph
	Relations []*relation
}

type relation struct {
	to *node
	distance float64
}

func NewGraph(data [][3]string) *Graph {
	var newGraph Graph = Graph{}
	var isThere bool = false
	var thisNode *node
	var dis float64

	for _, x := range data{
		isThere = false
		for _, y := range newGraph.Nodes{
			if y.name == x[0]{
				isThere = true
				thisNode = y
			}
		}
		if !isThere {
			thisNode = newGraph.addNode(x[0])
		}
		dis, _ = strconv.ParseFloat(x[2], 64)
		thisNode.addRelation(x[1], dis)

		isThere = false
		for _, y := range newGraph.Nodes{
			if y.name == x[1]{
				isThere = true
				thisNode = y
			}
		}
		if !isThere {
			thisNode = newGraph.addNode(x[1])
		}
		dis, _ = strconv.ParseFloat(x[2], 64)
		thisNode.addRelation(x[0], dis)
	}

	return  &newGraph
}

func (graph *Graph) addNode(name string) *node {
	var newNode node = node{name: name, graph: &*graph}
	graph.Nodes = append(graph.Nodes, &newNode)
	return &newNode
}


func (tNode *node) addRelation (toNode string, distance float64 ) *relation {
	x := false
	var Rel relation
	for _, to := range tNode.graph.Nodes{
		if to.name == toNode{
			x = true
			Rel = relation{to: *&to, distance: distance}
			tNode.Relations = append(tNode.Relations, &Rel)
			break
		}
	}

	if !x {
		newToNode := *tNode.graph.addNode(toNode)
		Rel = relation{to: &newToNode, distance: distance}
		tNode.Relations = append(tNode.Relations, &Rel)
	}
	return &Rel
}
