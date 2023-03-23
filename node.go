package gotoken

import (
	"errors"
	"fmt"
	"strings"
)

const (
	NodeDisplayCount = 8
)

type Node struct {
	Value       string
	Type        NodeType
	Parent      *Node
	Children    []*Node
	Constraints map[string]string
}

func (node *Node) Init(value string, parent *Node) *Node {
	node.Value = value
	node.Parent = parent
	node.Children = make([]*Node, 0)
	node.Constraints = make(map[string]string, 0)

	if value == "*" {
		node.Type = NodeTypeWildcard
	} else {
		node.Type = NodeTypeRegular
	}

	return node
}

func (node *Node) IsRoot() bool {
	return node.Parent == nil
}

func (node *Node) IsChild() bool {
	return node.Parent != nil
}

func (node *Node) IsWildcard() bool {
	return node.Type == NodeTypeWildcard
}

func (node *Node) HasWildcard() bool {
	for _, child := range node.Children {
		if child.IsWildcard() {
			return true
		}
	}

	return false
}

func (node *Node) SetParent(other *Node) {
	if node.HasParent() {
		node.Parent.RemoveChild(node)
		node.Parent = nil
	}

	node.Parent = other
	node.Parent.AddChild(node)
}

func (node *Node) HasParent() bool {
	return node.Parent != nil
}

func (node *Node) RemoveChild(other *Node) {
	children := make([]*Node, 0)

	for _, child := range node.Children {
		if child.Value == other.Value {
			continue
		}

		children = append(children, child)
	}

	node.Children = children
}

func (node *Node) HasChildren() bool {
	if node.Children == nil {
		return false
	}

	if len(node.Children) < 1 {
		return false
	}

	return true
}

func (node *Node) HasChild(value string, constraints map[string]string) bool {
	for _, child := range node.Children {
		if child.Value != value {
			continue
		}

		if !child.HasConstraints(constraints) {
			continue
		}

		return true
	}

	return false
}

func (node *Node) HasChildNode(other *Node) bool {
	return node.HasChild(other.Value, other.Constraints)
}

func (node *Node) GetChild(value string) *Node {
	if !node.HasChildren() {
		return nil
	}

	for _, child := range node.Children {
		if child.Value == value {
			return child
		}
	}

	return nil
}

func (node *Node) AddChild(other *Node) *Node {
	if other == nil {
		return nil
	}

	for _, child := range node.Children {
		if child.Value == other.Value {
			child.AddConstraints(other.Constraints)
			return other
		}
	}

	node.Children = append(node.Children, other)

	return other
}

func (node *Node) AddChildren(children ...*Node) {
	var (
		flag bool
	)

	if children == nil {
		return
	}

	if len(children) < 1 {
		return
	}

	if node.Children == nil {
		node.Children = make([]*Node, 0)
	}

	for _, child := range children {
		flag = true
		for _, v := range node.Children {
			if v.Value == child.Value {
				v.AddConstraints(child.Constraints)
				flag = false
				break
			}
		}

		if !flag {
			continue
		}

		node.Children = append(node.Children, child)
	}
}

func (node *Node) AddConstraints(constraints map[string]string) *Node {
	if node.Constraints == nil {
		node.Constraints = map[string]string{}
	}

	for k, v := range constraints {
		node.Constraints[k] = v
	}

	return node
}

func (node *Node) HasConstraints(constraints map[string]string) bool {
	for k, v := range constraints {
		if node.Constraints[k] != v {
			return false
		}
	}

	return true
}

func (node *Node) HasConstraint(key string, value string) bool {
	return node.Constraints[key] == value
}

func (node *Node) display_iter(space int) {
	space += NodeDisplayCount

	l := len(node.Children)
	i := 0

	// for ; i < (l / 2); i++ {
	// 	node.Children[i].display_iter(space)
	// }

	fmt.Printf(
		"%s#%s@%v\n",
		strings.Repeat(".", space),
		node.Value,
		node.Constraints)

	for ; i < l; i++ {
		node.Children[i].display_iter(space)
	}
}

func (node *Node) Display() {
	node.display_iter(-NodeDisplayCount)
}

func (node *Node) Includes(other *Node) bool {
	if node.IsWildcard() {
		return true
	}

	if other.Value != node.Value {
		return false
	}

	for _, otherChild := range other.Children {
		if node.IsWildcard() {
			continue
		}

		if node.HasWildcard() {
			continue
		}

		if !node.HasChildNode(otherChild) {
			return false
		}

		if !node.GetChild(otherChild.Value).Includes(otherChild) {
			return false
		}
	}

	return true
}

func NewNode(value string, parent ...*Node) (node *Node) {
	var (
		final *Node
	)

	final = nil
	if len(parent) >= 1 {
		final = parent[0]
	}

	node = (&Node{}).Init(value, final)
	if node.HasParent() {
		node.Parent.AddChild(node)
	}

	return
}

func NewTree(flat string) (root *Node, err error) {
	var (
		curr *Node
		next *Node
	)

	flat, _ = strings.CutPrefix(flat, "~.")

	root = NewNode("~")
	curr = root
	next = nil
	for _, nodeString := range strings.Split(flat, ".") {
		if nodeString == "" {
			continue
		}

		nodeSplit := strings.Split(nodeString, "@")
		nodeValue := nodeSplit[0]
		nodeConst := map[string]string{}

		if len(nodeSplit) >= 2 {
			for _, nodeConstString := range strings.Split(nodeSplit[1], ";") {
				nodeConstKeyVal := strings.Split(nodeConstString, "=")
				if len(nodeConstKeyVal) < 2 {
					err = errors.New("failed at parsing constraints")
					return
				}

				nodeConst[nodeConstKeyVal[0]] = nodeConstKeyVal[1]
			}
		}

		next = NewNode(nodeValue, curr)
		next.AddConstraints(nodeConst)
		curr.AddChildren(next)

		curr = next
		next = nil
	}

	return
}
