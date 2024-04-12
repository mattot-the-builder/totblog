package auth_test

import (
	"testing"

	"github.com/mattot-the-builder/totblog/internal/auth"
)

func TestGenerateToken(t *testing.T) {
	token, err := auth.GenerateToken("test")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(token)
}

func TestVerifyToken(t *testing.T) {
	token, err := auth.GenerateToken("test")
	if err != nil {
		t.Fatal(err)
	}
	err = auth.ValidateToken(token)
	if err != nil {
		t.Fatal(err)
	}
}
