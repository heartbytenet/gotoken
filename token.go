package gotoken

import "github.com/google/uuid"

type Token struct {
	UUID  string
	Perms map[string]struct{}
	Tree  *Node
}

func (token *Token) Init(uuid string) *Token {
	token.UUID = uuid
	token.Perms = map[string]struct{}{}
	token.Tree = NewNodeTree("~")

	return token
}

func (token *Token) AddPerms(perms ...string) {
	for _, perm := range perms {
		token.Perms[perm] = struct{}{}
		token.Tree.Merge(NewNodeTree(perm))
	}
}

func (token *Token) HasPerm(perm string) bool {
	return token.Tree.Includes(
		NewNodeTree(perm))
}

func NewToken() (token *Token) {
	token = (&Token{}).Init(uuid.New().String())

	return
}
