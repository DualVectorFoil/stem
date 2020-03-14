package encryptUtil

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"github.com/DualVectorFoil/stem/app/conf"
)

func EncryptPassword(password []byte) string {
	block, _ := pem.Decode([]byte(conf.ACCOUNT_SERVICE_PUB_KEY))
	if block == nil {
		panic("failed to parse PEM block containing the public key")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic("failed to parse DER encoded public key: " + err.Error())
	}

	pubKey := pub.(*rsa.PublicKey)

	cipherPassword, err := rsa.EncryptPKCS1v15(rand.Reader, pubKey, password)
	if err != nil {
		panic("failed to encrypt the password with PKCS1v15" + err.Error())
	}
	base64Password := base64.StdEncoding.EncodeToString(cipherPassword)

	return base64Password
}
