package auth

import (
	"testing"
)

func TestGetHashedPassword(t *testing.T) {
	pass := "passwordtest"
	hpass, err := GetHashedPassword(pass)
	t.Log(hpass, err)
}
