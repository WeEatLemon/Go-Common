package authenticator

import (
	"fmt"
	"testing"
)

func Test_InitAuth(t *testing.T) {
	secret, code := initAuth("test@test.com", "test")
	fmt.Println("secret", secret)
	fmt.Println("code", code)
}

//Secret: 2TFQHJ32DXDVASGXZ3I3SE67THJLQP4N
//Code: 093050 <nil>
//Qrcode otpauth://totp/test@test.com?secret=093050
//QrcodeUrl https://api.qrserver.com/v1/create-qr-code/?data=otpauth%3A%2F%2Ftotp%2Ftest%40test.com%3Fsecret%3D2TFQHJ32DXDVASGXZ3I3SE67THJLQP4N&size=200x200&ecc=M
func TestGoogleAuth_VerifyCode(t *testing.T) {
	b, err := NewGoogleAuth().VerifyCode("2TFQHJ32DXDVASGXZ3I3SE67THJLQP4N", "644475")
	if err != nil {
		fmt.Println("err", err)
	} else {
		fmt.Println("bool", b)
	}
}
