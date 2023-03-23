package gotoken

import "log"

type Token struct {
	UUID  string
	Perms map[string]*Node
}

func (token *Token) Init(uuid string) *Token {
	token.UUID = uuid
	token.Perms = map[string]*Node{}

	return token
}

func (token *Token) AddPerms(perms ...string) (res *Token, err error) {
	var (
		node *Node
	)

	res = token

	for _, perm := range perms {
		if token.Perms[perm] != nil {
			continue
		}

		node, err = NewTree(perm)
		if err != nil {
			return
		}

		token.Perms[perm] = node
	}

	return
}

func (token *Token) AddPermsMust(perms ...string) *Token {
	var (
		err error
	)

	_, err = token.AddPerms(perms...)
	if err != nil {
		log.Fatalln("error while adding perms", err)
	}

	return token
}

func (token *Token) HasPerm(perm string) bool {
	var (
		node *Node
		err  error
	)

	node, err = NewTree(perm)
	if err != nil {
		return false
	}

	for key, val := range token.Perms {
		if key == perm {
			return true
		}

		if val.Includes(node) {
			return true
		}
	}

	return false
}
