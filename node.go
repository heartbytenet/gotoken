package gotoken

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

const (
	NodeDisplayCount = 16
)

type Node struct {
	uuid        uuid.UUID
	Type        NodeType
	Parent      *Node
	Children    []*Node
	Constraints map[string]string
}

func (node *Node) Init(parent *Node) *Node {
	node.uuid = uuid.New()

	node.Parent = parent
	node.Children = make([]*Node, 0)
	node.Constraints = make(map[string]string, 0)

	return node
}

func (node *Node) IsRoot() bool {
	return node.Parent == nil
}

func (node *Node) IsChild() bool {
	return node.Parent != nil
}

func (node *Node) HasParent() bool {
	return node.Parent != nil
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
			if v.uuid == child.uuid {
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

func (node *Node) HasConstraint(value string) bool {
	for _, v := range node.Constraints {
		if v == value {
			return true
		}
	}

	return false
}

func (node *Node) display_iter(space int) {
	space += NodeDisplayCount

	l := len(node.Children)
	i := 0

	// for ; i < (l / 2); i++ {
	// 	node.Children[i].display_iter(space)
	// }

	fmt.Println(strings.Repeat("-", space) + node.uuid.String()[:8])

	for ; i < l; i++ {
		node.Children[i].display_iter(space)
	}
}

func (node *Node) Display() {
	node.display_iter(-NodeDisplayCount)
}

func (node *Node) Has(other *Node) bool {
	return false
}

func NewNode(parent ...*Node) *Node {
	var (
		final *Node
	)

	final = nil
	if len(parent) >= 1 {
		final = parent[0]
	}

	return (&Node{}).Init(final)
}
