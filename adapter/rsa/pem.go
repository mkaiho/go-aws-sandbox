package rsa

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"

	"golang.org/x/crypto/ssh"
)

type PemBlockType string

func (t PemBlockType) String() string {
	return string(t)
}

const (
	PemBlockTypePKCS1PrivateKey   PemBlockType = "RSA PRIVATE KEY"
	PemBlockTypePKCS8PrivateKey   PemBlockType = "PRIVATE KEY"
	PemBlockTypeOpenSSHPrivateKey PemBlockType = "OPENSSH PRIVATE KEY"
)

type PemManager struct{}

func NewPemManager() *PemManager {
	return &PemManager{}
}

func (m *PemManager) ReadRSAPrivateKey(path string) (*rsa.PrivateKey, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(bytes)
	if block == nil {
		return nil, fmt.Errorf("pem block is empty")
	}

	var key *rsa.PrivateKey
	switch blockType := (PemBlockType)(block.Type); blockType {
	case PemBlockTypePKCS1PrivateKey:
		key, err = m.parsePKCS1PrivateBlock(block.Bytes)
	case PemBlockTypePKCS8PrivateKey:
		key, err = m.parsePKCS8PrivateBlock(block.Bytes)
	case PemBlockTypeOpenSSHPrivateKey:
		key, err = m.parseOpenSSHPrivateBlock(bytes)
	default:
		return nil, fmt.Errorf("invalid block type: %s", blockType)
	}
	if err != nil {
		return nil, err
	}

	key.Precompute()

	if err := key.Validate(); err != nil {
		return nil, err
	}

	return key, nil
}

func (m *PemManager) parsePKCS1PrivateBlock(der []byte) (*rsa.PrivateKey, error) {
	return x509.ParsePKCS1PrivateKey(der)
}

func (m *PemManager) parsePKCS8PrivateBlock(der []byte) (*rsa.PrivateKey, error) {
	parsedKey, err := x509.ParsePKCS8PrivateKey(der)
	if err != nil {
		return nil, err
	}
	rsaKey, ok := parsedKey.(*rsa.PrivateKey)
	if !ok {
		return nil, fmt.Errorf("not RSA private key")
	}
	return rsaKey, nil
}

func (m *PemManager) parseOpenSSHPrivateBlock(pemBytes []byte) (*rsa.PrivateKey, error) {
	parsedKey, err := ssh.ParseRawPrivateKey(pemBytes)
	if err != nil {
		return nil, err
	}
	rsaKey, ok := parsedKey.(*rsa.PrivateKey)
	if !ok {
		return nil, fmt.Errorf("not RSA private key")
	}
	return rsaKey, nil
}
