package rsa

import (
	"fmt"
	"github.com/IEatLemons/GoBase/common/helper"
	"github.com/IEatLemons/GoBase/config"
	"github.com/yuchenfw/gocrypt"
	"testing"
)

type RSACryptTest struct {
	data       string
	encodeType gocrypt.Encode
	hashType   gocrypt.Hash
}

var (
	rsaTest = RSACryptTest{}
)

func Init() {
	vHandle := config.New()
	path, _ := helper.GetProjectRoot()
	path = path + "/../../../"
	err := vHandle.InitConfig(path)
	if err != nil {

	}
	InitTRSACrypt(vHandle.Config.RMRsa)
}

func TestRsaCrypt_Decrypt(t *testing.T) {
	Init()
	RSASecret := NewRSACrypt()

	rsaTest.data = "test1"
	rsaTest.encodeType = gocrypt.Base64
	rsaTest.hashType = gocrypt.SHA256

	//encrypt data & encode result
	encrypt, err := RSASecret.Encrypt(rsaTest.data, rsaTest.encodeType)

	if err != nil {
		t.Fatalf("encrypt error : %v", err)
	}
	fmt.Println("encrypt", encrypt)
}

func BenchmarkRSACrypt(t *testing.B) {
	Init()
	handle := NewRSACrypt()
	var rsaTests = []RSACryptTest{
		{
			"test",
			gocrypt.HEX,
			gocrypt.MD5,
		},
		{
			"base64",
			gocrypt.Base64,
			gocrypt.SHA256,
		},
		{
			"01234567890123456789012345678901234567890123456789012",
			gocrypt.String,
			gocrypt.SHA1,
		},
	}
	for _, rsaTest := range rsaTests {
		fmt.Println(" ================================== ")
		//encrypt data & encode result
		fmt.Printf("data is %+v \n", rsaTest.data)
		encrypt, err := handle.Encrypt(rsaTest.data, rsaTest.encodeType)
		if err != nil {
			t.Fatalf("encrypt error : %v", err)
		}
		fmt.Printf("encrypt is %+v \n", encrypt)
		//decrypt encrypted & encoded data
		decrypt, err := handle.Decrypt(encrypt, rsaTest.encodeType)
		if err != nil {
			t.Fatalf("decrypt error : %v", err)
		}
		fmt.Printf("decrypt is %+v \n", decrypt)
		if decrypt != rsaTest.data {
			t.Fatalf("decrypt get result %s , want get %s ", decrypt, rsaTest.data)
		}
		//sign data with digest algorithm & encode result
		sign, err := handle.Sign(rsaTest.data, rsaTest.hashType, rsaTest.encodeType)
		if err != nil {
			t.Fatalf("sign error : %v", err)
		}
		fmt.Printf("sign is %+v \n", sign)
		//verify data that signed with digest algorithm & encoded whether match original data
		verifySign, err := handle.VerifySign(rsaTest.data, rsaTest.hashType, sign, rsaTest.encodeType)
		if err != nil {
			t.Fatalf("verifySign error : %v", err)
		}
		fmt.Printf("verifySign is %+v \n", verifySign)
	}
}
