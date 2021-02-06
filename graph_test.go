package shortpath

import (
	"fmt"
	"testing"
)

func TestNewGraph(t *testing.T) {
	x := [][3]string{{"A","B","10"},{"A","C","20"},{"B","C","2"}}
	g := NewGraph(x)
	for _,i := range g.Nodes{
		fmt.Println(i.name)
		for _,r := range i.Relations{
			fmt.Printf("to: %v - distance: %v", r.to, r.distance)
		}
	}
}