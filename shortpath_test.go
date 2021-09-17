package shortpath

import (
	"fmt"
	"testing"
)

func TestShortPath(t *testing.T)  {
	x := [][3]string{{"A","B","3"},{"A","C","2"},{"A","D","9"},{"B","E","3"},{"C","D","5"},{"D","E","2"},{"E","F","1"}}
	g := NewGraph(x)
	n := g.Nodes[0]
	p := n.toAllNode()



	fmt.Println(p)

}