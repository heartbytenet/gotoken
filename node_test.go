package gotoken_test

import (
	"testing"

	"github.com/heartbytenet/gotoken"
)

func TestNodeNew(t *testing.T) {
	node := gotoken.NewNode("eulav")
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
	parent := gotoken.NewNode("this_is_parent")
	if parent.HasParent() {
		t.Error("parent node should not have a parent")
		return
	}

	if parent.Parent != nil {
		t.Error("parent node should not have a parent")
		return
	}

	if parent.HasChildren() {
		t.Error("parent node should not have children for now")
		return
	}

	childX := gotoken.NewNode("this_is_childX", parent)
	childY := gotoken.NewNode("this_is_childY", nil)
	childZ := gotoken.NewNode("this_is_childZ")
	childW := gotoken.NewNode("this_is_childW")

	childW.Parent = parent

	if !childX.HasParent() {
		t.Error("childX node should have a parent")
		return
	}

	if childX.Parent == nil {
		t.Error("childX node should have a parent")
		return
	}

	if childY.HasParent() {
		t.Error("childY node should not have a parent")
		return
	}

	if childY.Parent != nil {
		t.Error("childY node should not have a parent")
		return
	}

	if childZ.HasParent() {
		t.Error("childZ node should not have a parent")
		return
	}

	if childZ.Parent != nil {
		t.Error("childZ node should not have a parent")
		return
	}

	if !childW.HasParent() {
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

func TestTreeNewEmpty(t *testing.T) {
	_, err := gotoken.NewTree("")
	if err != nil {
		t.Error(err)
		return
	}
}

func TestTreeNewUnnamed(t *testing.T) {
	root, err := gotoken.NewTree("")
	if err != nil {
		t.Error(err)
		return
	}

	if root.Value != "~" {
		t.Error("root value should always be ~")
		return
	}

	if root.HasChildren() {
		t.Error("root value should not have children now")
		return
	}
}

func TestTreeNew(t *testing.T) {
	root, err := gotoken.NewTree("hello")
	if err != nil {
		t.Error(err)
		return
	}

	if root.Value != "~" {
		t.Error("root value should always be ~")
		return
	}

	if !root.HasChildren() {
		t.Error("root value should have one child")
		return
	}

	node := root.GetChild("hello")
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

	if !node.HasParent() {
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
		err   error
	)

	treeX, err = gotoken.NewTree("hello.world.I.am.a.tree.that.is.very.special")
	if err != nil {
		t.Error(err)
		return
	}

	treeY, err = gotoken.NewTree("hello.world.I.am.a.tree")
	if err != nil {
		t.Error(err)
		return
	}

	treeZ, err = gotoken.NewTree("hello.world.I.am.an.imposter")
	if err != nil {
		t.Error(err)
		return
	}

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
		err   error
	)

	treeX, err = gotoken.NewTree("hello.sekai")
	if err != nil {
		t.Error(err)
		return
	}

	treeY, err = gotoken.NewTree("*")
	if err != nil {
		t.Error(err)
		return
	}

	treeZ, err = gotoken.NewTree("hello.*")
	if err != nil {
		t.Error(err)
		return
	}

	if !treeX.Includes(treeY) {
		t.Error("treeX should include treeY")
		return
	}

	if !treeX.Includes(treeZ) {
		t.Error("treeX should include treeZ")
		return
	}

	if !treeZ.Includes(treeY) {
		t.Error("treeZ should include treeY")
		return
	}
}
