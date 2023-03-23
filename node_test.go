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

	treeX, err = gotoken.NewTree("*")
	if err != nil {
		t.Error(err)
		return
	}

	treeY, err = gotoken.NewTree("hello.*")
	if err != nil {
		t.Error(err)
		return
	}

	treeZ, err = gotoken.NewTree("hello.sekai")
	if err != nil {
		t.Error(err)
		return
	}

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

func TestToken_Init(t *testing.T) {
	uuid := "ce93fe81-5cf5-4819-b898-5f28c640bedd"
	token := (&gotoken.Token{}).Init(uuid)

	if token == nil {
		t.Error()
		return
	}

	if token.UUID != uuid {
		t.Error()
		return
	}
}

func TestToken_AddPerms(t *testing.T) {
	var (
		err error
	)

	token := (&gotoken.Token{}).Init("ce93fe81-5cf5-4819-b898-5f28c640bedd")

	_, err = token.AddPerms("hello.world")
	if err != nil {
		t.Error(err)
		return
	}

	token.AddPermsMust("hello.world@x=y")
}

func TestToken_HasPerm(t *testing.T) {
	token := (&gotoken.Token{}).Init("ce93fe81-5cf5-4819-b898-5f28c640bedd")

	token.AddPermsMust("hello.sekai")
	token.AddPermsMust("hello.world@abc=123")
	token.AddPermsMust("hello.world@xyz=987.hey")

	if token.HasPerm("*") {
		t.Error()
	}

	if !token.HasPerm("hello.sekai") {
		t.Error()
	}

	if !token.HasPerm("hello.world") {
		t.Error()
	}

	if token.HasPerm("hello.world@123=abc") {
		t.Error()
	}

	if !token.HasPerm("hello.world@abc=123") {
		t.Error()
	}

	if !token.HasPerm("hello.world@xyz=987") {
		t.Error()
	}

	if !token.HasPerm("hello.world@xyz=987.hey") {
		t.Error()
	}

	if token.HasPerm("hello.world@abc=123.hey") {
		t.Error()
	}
}
