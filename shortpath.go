package shortpath

var infinity float64 = 99999999

type shortPaths struct {
	from *node
	to map[string]float64
}

func (fn *node) toAllNode() shortPaths{
	unvisited := make([]*node, len(fn.graph.Nodes))
	copy(unvisited, fn.graph.Nodes)
	var toVisit []*node
	sp := shortPaths{from: fn, to: make(map[string]float64)}

	ifn, err := getNodeIndex(unvisited, fn)
	if err != nil {
		panic(err)
	}
	unvisited = append(unvisited[:ifn], unvisited[ifn+1:]...)

	for _,j := range fn.graph.Nodes {
		if fn.name == j.name{
			sp.to[j.name] = 0
		}else {
			sp.to[j.name] = infinity
		}
	}

	for _,r := range fn.Relations {
		sp.to[r.to.name] = r.distance
		ind, err := getNodeIndex(fn.graph.Nodes, r.to)
		if err != nil {
			panic(err)
		}
		toVisit = append(toVisit, fn.graph.Nodes[ind])
	}

	var currentNode *node
	var distance float64

	for len(unvisited) != 0 {

		in, nil := getNodeIndex(fn.graph.Nodes, toVisit[0])
		if err != nil {
			panic(err)
		}
		currentNode = fn.graph.Nodes[in]
		distance = sp.to[currentNode.name]

		for _, rel := range currentNode.Relations {
			if distance + rel.distance < sp.to[rel.to.name] {
				sp.to[rel.to.name] = distance + rel.distance
			}
			c := !isInclude(toVisit, rel.to) && isInclude(unvisited, rel.to)
			if  c {
				toVisit = append(toVisit, rel.to)
			}
		}

		i, nil := getNodeIndex(unvisited, currentNode)
		if err != nil {
			panic(err)
		}
		unvisited = append(unvisited[:i], unvisited[i+1:]...)

		toVisit = toVisit[1:]

	}

	return sp
}