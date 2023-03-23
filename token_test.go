package gotoken_test

import (
	"github.com/heartbytenet/gotoken"
	"testing"
)

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
