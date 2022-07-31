package ca

import (
	"crypto/rand"
	"crypto/rsa"
	cx509 "crypto/x509"
	"crypto/x509/pkix"
	"log"

	"math/big"
	mathRand "math/rand"
	"time"
)

var CA CertificateAuthority

type CertificateAuthority struct {
	RootCA		*cx509.Certificate
	PrivateKey 	*rsa.PrivateKey
}


func (ca *CertificateAuthority) makeRootCA() error {
	privateKey, err :=rsa.GenerateKey(rand.Reader,2048)
	if err !=nil {
		log.Print("error happens when generate private key to create root CA")
		return err
	}
	mathRand.Seed(time.Now().UnixNano())
	rootCertificateTemplate := cx509.Certificate{
		Version:      1,
		SerialNumber: big.NewInt((int64)(mathRand.Int())),
		Subject: pkix.Name{
			Country:            []string{"CN"},
			Organization:       []string{"Fudan"},
			OrganizationalUnit: []string{"Mathematics"},
			Locality:           []string{"上海"},
			Province:           []string{"Shanghai"},
			StreetAddress:      []string{"Handan Road #200"},
			PostalCode:         []string{"200201"},
			CommonName:         "Fudan CA",
		},

		EmailAddresses: []string{"jacky01.zhang@outlook.com"},
		DNSNames:       []string{"localhost"},

		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(10, 0, 0),
		IsCA:                  true,
		BasicConstraintsValid: true,
	}
	cx509.CreateCertificate(rand.Reader,&rootCertificateTemplate,&rootCertificateTemplate,&privateKey.PublicKey,privateKey)
	return nil
}
