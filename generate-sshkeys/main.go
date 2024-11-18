package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"

	"golang.org/x/crypto/ssh"
)

func main() {
	private_key, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		fmt.Println("err:", err.Error())
		return
	}

	private_key_pem := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(private_key),
	})
	if private_key_pem == nil {
		fmt.Println("err: cannot be encoded")
		return
	}

	privateKeyPath := "id_rsa" // Replace with your desired file path
	err = os.WriteFile(privateKeyPath, private_key_pem, 0600)
	if err != nil {
		fmt.Printf("Failed to save private key: %v\n", err)
		return
	}

	public_key, err := ssh.NewPublicKey(&private_key.PublicKey)
	if err != nil {
		fmt.Println("err in public_key:", err.Error())
		return
	}

	public_key_in_bytes := ssh.MarshalAuthorizedKey(public_key)
	err = os.WriteFile("id_rsa.pub", public_key_in_bytes, 0644)
	if err != nil {
		fmt.Printf("Failed to save public key: %v\n", err)
		return
	}

	fmt.Println("it is to generate", private_key_pem)
}
