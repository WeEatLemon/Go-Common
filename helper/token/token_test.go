package token

import (
	"fmt"
	"testing"
)

func TestRandStringRunes(t *testing.T) {
	token := RandStringRunes(32)

	fmt.Printf("%+v \n", token)
}
