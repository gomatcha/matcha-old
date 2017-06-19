package radix

import (
	"fmt"
	"strconv"
	"strings"
)

type Radix struct {
	root *Node
}

func NewRadix() *Radix {
	return &Radix{
		root: &Node{},
	}
}

func (r *Radix) Insert(path []int64) *Node {
	return r.root.insert(path)
}

func (r *Radix) At(path []int64) *Node {
	return r.root.at(path)
}

func (r *Radix) Delete(path []int64) {
	// Cannot delete root node.
	if len(path) == 0 {
		return
	}
	r.root.delete(path)
}

func (r *Radix) String() string {
	return r.root.debugString()
}

type Node struct {
	children map[int64]*Node
	exists   bool
	Value    interface{}
}

func (n *Node) insert(path []int64) *Node {
	if len(path) == 0 {
		n.exists = true
		return n
	}

	child, ok := n.children[path[0]]
	if !ok {
		child = &Node{}
		if n.children == nil {
			n.children = map[int64]*Node{}
		}
		n.children[path[0]] = child
	}
	return child.insert(path[1:])
}

func (n *Node) at(path []int64) *Node {
	if len(path) == 0 {
		if n.exists == false {
			return nil
		}
		return n
	}
	child, ok := n.children[path[0]]
	if !ok {
		return nil
	}
	return child.at(path[1:])
}

func (n *Node) delete(path []int64) {
	if len(path) == 1 {
		child, ok := n.children[path[0]]
		if !ok {
			return
		}

		if len(child.children) == 0 {
			n.Value = nil
			delete(n.children, path[0])
		} else {
			child.exists = false
		}
		return
	}

	child, ok := n.children[path[0]]
	if !ok {
		return
	}
	child.delete(path[1:])
}

func (n *Node) debugString() string {
	all := []string{}
	for k, i := range n.children {
		lines := strings.Split(i.debugString(), "\n")
		for idx, line := range lines {
			if idx == 0 {
				lines[idx] = padRight(strconv.Itoa(int(k)), " ", 5) + line
			} else {
				lines[idx] = "|    " + line
			}
		}
		all = append(all, lines...)
	}

	str := fmt.Sprintf("{%p Exists:%v Value:%v}", n, n.exists, n.Value)
	if len(all) > 0 {
		str += "\n" + strings.Join(all, "\n")
	}
	return str
}

func padRight(str, pad string, length int) string {
	for {
		str += pad
		if len(str) > length {
			return str[0:length]
		}
	}
}
