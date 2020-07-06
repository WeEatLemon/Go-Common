package password

import (
	"fmt"
	"testing"
)

func TestHashPassword(t *testing.T) {
	str := "e10adc3949ba59abbe56e057f20f883e"

	pwd, err := HashPassword(str)

	fmt.Printf("err : %+v \n", err)
	fmt.Printf("pwd : %+v \n", pwd)
}

func TestCheckPasswordHash(t *testing.T) {
	str := "123456"
	hash := "$2a$14$3.e6oPek5P/idO0E1g3lgeit0N0bYy1REMkb/THUEmoTGI6G1FGN2"

	err := CheckPasswordHash(str, hash)
	fmt.Printf("err : %+v \n", err)
}
