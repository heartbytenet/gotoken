package gotoken_test

import (
	"testing"

	"github.com/heartbytenet/gotoken"
)

func TestNodeNew(t *testing.T) {
	node := gotoken.NewNode("eulav", nil, nil)
	if node == nil {
		t.Error("node is nil")
		return
	}

	if node.Value != "eulav" {
		t.Error("node has incorrect value")
		return
	}
}

func TestNodeNewParent(t *testing.T) {
	parent := gotoken.NewNode("this_is_parent", nil, nil)
	if parent.IsChild() {
		t.Error("parent node should not have a parent")
		return
	}

	if parent.Parent != nil {
		t.Error("parent node should not have a parent")
		return
	}

	if parent.IsParent() {
		t.Error("parent node should not have children for now")
		return
	}

	childX := gotoken.NewNode("this_is_childX", nil, parent)
	childY := gotoken.NewNode("this_is_childY", nil, nil)
	childZ := gotoken.NewNode("this_is_childZ", nil, nil)
	childW := gotoken.NewNode("this_is_childW", nil, nil)

	childW.Parent = parent

	if !childX.IsChild() {
		t.Error("childX node should have a parent")
		return
	}

	if childX.Parent == nil {
		t.Error("childX node should have a parent")
		return
	}

	if childY.IsChild() {
		t.Error("childY node should not have a parent")
		return
	}

	if childY.Parent != nil {
		t.Error("childY node should not have a parent")
		return
	}

	if childZ.IsChild() {
		t.Error("childZ node should not have a parent")
		return
	}

	if childZ.Parent != nil {
		t.Error("childZ node should not have a parent")
		return
	}

	if !childW.IsChild() {
		t.Error("childW node should have a parent")
		return
	}

	if childW.Parent == nil {
		t.Error("childW node should have a parent")
		return
	}

	childZ.SetParent(childX)
	if childZ.Parent == nil || childZ.Parent != childX {
		t.Error("childZ node should have childX node as parent")
		return
	}

	if childZ.Parent.Parent != parent {
		t.Error("childZ node should have parent node as parent parent")
		return
	}
}

func TestTreeNewUnnamed(t *testing.T) {
	root := gotoken.NewNodeTree("")

	if root.Value != "~" {
		t.Error("root value should always be ~")
		return
	}

	if root.IsParent() {
		t.Error("root value should not have children now")
		return
	}
}

func TestTreeNew(t *testing.T) {
	root := gotoken.NewNodeTree("hello")

	if root.Value != "~" {
		t.Error("root value should always be ~")
		return
	}

	if len(root.Children) > 1 {
		t.Error("root value should have one child")
		return
	}

	node := root.Child("hello", nil)
	if node == nil {
		t.Error("child node should not be nil")
		return
	}

	if node.Value != "hello" {
		t.Error("child node value is invalid")
		return
	}

	if node.Parent == nil {
		t.Error("chid node should have a parent")
		return
	}

	if !node.IsChild() {
		t.Error("child node should have a parent")
		return
	}

	if node.Parent != root {
		t.Error("child node should have root node as parent")
		return
	}
}

func TestTreeIncludesSimple(t *testing.T) {
	var (
		treeX *gotoken.Node
		treeY *gotoken.Node
		treeZ *gotoken.Node
	)

	treeX = gotoken.NewNodeTree("hello.world.I.am.a.tree.that.is.very.special")
	treeY = gotoken.NewNodeTree("hello.world.I.am.a.tree")
	treeZ = gotoken.NewNodeTree("hello.world.I.am.an.imposter")

	if !treeX.Includes(treeX) {
		t.Error("a tree should always include itself")
		return
	}

	if !treeY.Includes(treeY) {
		t.Error("a tree should always include itself")
		return
	}

	if !treeZ.Includes(treeZ) {
		t.Error("a tree should always include itself")
		return
	}

	if !treeX.Includes(treeY) {
		t.Error()
		return
	}

	if treeX.Includes(treeZ) {
		t.Error()
		return
	}

	if treeY.Includes(treeX) {
		t.Error()
		return
	}

	if treeY.Includes(treeZ) {
		t.Error()
		return
	}

	if treeZ.Includes(treeX) {
		t.Error()
		return
	}

	if treeZ.Includes(treeY) {
		t.Error()
		return
	}
}

func TestTreeIncludeWildcard(t *testing.T) {
	var (
		treeX *gotoken.Node
		treeY *gotoken.Node
		treeZ *gotoken.Node
	)

	treeX = gotoken.NewNodeTree("*")
	treeY = gotoken.NewNodeTree("hello.*")
	treeZ = gotoken.NewNodeTree("hello.sekai")

	if !treeX.Includes(treeX) {
		t.Error("trees always include themselves")
		return
	}

	if !treeY.Includes(treeY) {
		t.Error("trees always include themselves")
		return
	}

	if !treeZ.Includes(treeZ) {
		t.Error("trees always include themselves")
		return
	}

	if !treeX.Includes(treeX) {
		t.Error()
		return
	}

	if !treeX.Includes(treeY) {
		t.Error()
		return
	}

	if !treeX.Includes(treeZ) {
		t.Error()
		return
	}

	if !treeY.Includes(treeZ) {
		t.Error()
		return
	}
}
