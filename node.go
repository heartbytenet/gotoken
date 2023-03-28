package gotoken

import (
	"fmt"
	"strings"
)

const (
	NodeDisplayCount = 8

	NodePrefixRoot     = "~"
	NodePrefixWildcard = "*"
)

type Node struct {
	Parent   *Node
	Children map[*Node]struct{}

	Type NodeType

	Value       string
	Constraints map[string]string
}

func getFirstChar(v string) string {
	if len(v) < 1 {
		return ""
	}
	return v[0:1]
}

func (node *Node) IsWildcard() bool {
	return node.Value == NodePrefixWildcard
}

func (node *Node) IsChild() bool {
	return node.Parent != nil
}

func (node *Node) IsParent() bool {
	return len(node.Children) >= 1
}

func (node *Node) SetParent(other *Node) {
	if other == nil {
		return
	}

	if node.Parent != nil {
		delete(node.Parent.Children, node)
	}

	node.Parent = other
	other.Children[node] = struct{}{}
}

func (node *Node) Init(value string, constraints map[string]string, parent *Node) *Node {
	node.Parent = parent
	node.Children = make(map[*Node]struct{}, 16)

	node.Value = value
	node.Constraints = constraints
	if node.Constraints == nil {
		node.Constraints = make(map[string]string)
	}

	switch getFirstChar(value) {
	case NodePrefixWildcard:
		node.Type = NodeTypeWildcard
		break
	default:
		node.Type = NodeTypeRegular
		break
	}

	return node
}

func (node *Node) Match(value string, constraints map[string]string) bool {
	if node.Value != value {
		return false
	}

	if len(constraints) != len(node.Constraints) {
		return false
	}

	for key, val := range constraints {
		if node.Constraints[key] != val {
			return false
		}
	}

	return true
}

func (node *Node) MatchWith(value string, constraints map[string]string) bool {
	if node.Value != value {
		return false
	}

	for key, val := range constraints {
		if node.Constraints[key] != val {
			return false
		}
	}

	return true
}

func (node *Node) Child(value string, constraints map[string]string) *Node {
	var (
		wild *Node
	)

	wild = nil

	for child := range node.Children {
		if child.IsWildcard() {
			wild = child
			continue
		}

		if !child.Match(value, constraints) {
			continue
		}

		return child
	}

	return wild // Todo: I'm not sure about this, the idea is to return a wildcard at the end if there's one anyway
}

func (node *Node) ChildWith(value string, constraints map[string]string) *Node {
	var (
		wild *Node
	)

	wild = nil

	for child := range node.Children {
		if child.IsWildcard() {
			wild = child
			continue
		}

		if !child.MatchWith(value, constraints) {
			continue
		}

		return child
	}

	return wild // Todo: same thing here as node.Child()
}

func (node *Node) Append(value string, constraints map[string]string) (other *Node) {
	other = node.Child(value, constraints)
	if other != nil {
		return
	}

	other = NewNode(value, constraints, node)
	return
}

func (node *Node) Display(_space ...int) {
	space := 0
	if len(_space) >= 1 {
		space = _space[0]
	}

	fmt.Printf(
		"%s#%s@%v\n",
		strings.Repeat(".", space),
		node.Value,
		node.Constraints)

	for child := range node.Children {
		child.Display(space + NodeDisplayCount)
	}
}

func (node *Node) Includes(other *Node) bool {
	var (
		child *Node
	)

	if node.IsWildcard() {
		return true
	}

	if !node.MatchWith(other.Value, other.Constraints) {
		return false
	}

	for n := range other.Children {
		child = node.ChildWith(n.Value, n.Constraints)
		if child == nil {
			return false
		}

		if !child.Includes(n) {
			return false
		}
	}

	return true
}

func (node *Node) Merge(other *Node) {
	if !node.Match(other.Value, other.Constraints) {
		return
	}

	for child := range other.Children {
		node.Append(child.Value, child.Constraints).Merge(child)
	}
}

// NewNode creates a node and adds its reference to the parent if existing
func NewNode(value string, constraints map[string]string, parent *Node) (node *Node) {
	node = (&Node{}).Init(
		value, constraints, parent)

	if parent != nil {
		parent.Children[node] = struct{}{}
	}

	return
}

func NewNodeTree(source string) (root *Node) {
	var (
		curr *Node
		next *Node
	)

	root = NewNode(NodePrefixRoot, nil, nil)
	if source == "~" {
		return
	}

	curr = root
	next = nil

	for _, nodeSource := range strings.Split(source, ".") {
		if nodeSource == "" {
			continue
		}

		nodeSourceSplit := strings.Split(nodeSource, "@")

		nodeValue := nodeSourceSplit[0]
		if nodeValue == "" {
			continue
		}

		nodeConstraints := map[string]string{}
		if len(nodeSourceSplit) >= 2 {
			nodeConstraintsSource := nodeSourceSplit[1]
			if nodeConstraintsSource != "" {
				for _, nodeConstraint := range strings.Split(nodeConstraintsSource, ";") {
					if nodeConstraint == "" {
						continue
					}

					nodeConstraintSplit := strings.Split(nodeConstraint, "=")
					if len(nodeConstraintSplit) < 2 {
						continue
					}

					if nodeConstraintSplit[0] == "" {
						continue
					}

					if nodeConstraintSplit[1] == "" {
						continue
					}

					nodeConstraints[nodeConstraintSplit[0]] = nodeConstraintSplit[1]
				}
			}
		}

		next = curr.Append(nodeValue, nodeConstraints)

		curr = next
		next = nil
	}

	return
}
