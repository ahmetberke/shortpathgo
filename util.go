package shortpath

import "errors"

func isInclude(list []*node, node *node) bool {
	for _,n := range list {
		if n.name == node.name {
			return true
		}
	}
	return false
}

func getNodeIndex(list []*node, node *node) (int, error) {
	for i,v := range list {
		if v.name == node.name {
			return i, nil
		}
	}
	return 0, errors.New("not found this node")
}