package utils

import (
	"crypto/ed25519"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
)

// GenerateKeyPair generates an ed25519 keypair
func GenerateKeyPair() (ed25519.PublicKey, ed25519.PrivateKey, error) {
	return ed25519.GenerateKey(rand.Reader)
}

// GenerateKeyPairFromSeed generates an ed25519 keypair
func GenerateKeyPairFromSeed(seed []byte) (ed25519.PublicKey, ed25519.PrivateKey, error) {
	keypair := ed25519.NewKeyFromSeed(seed)

	pubKey, ok := keypair.Public().(ed25519.PublicKey)

	if !ok {
		return nil, nil, errors.New("public key is not of type ed25519")
	}

	return pubKey, keypair, nil
}

// PEMEncodePublicKey encodes an ed25519 public key using the common PKIX standard
// with a PEM format
func PEMEncodePublicKey(pubKey ed25519.PublicKey) (pub []byte, err error) {
	pubBytes, err := x509.MarshalPKIXPublicKey(pubKey)
	if err != nil {
		return nil, err
	}

	pemEncodedPubKey := pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pubBytes,
	})

	return pemEncodedPubKey, nil
}

// PEMEncodePrivateKey encodes an ed25519 private key using the common PKCS8 standard
// with a PEM format
func PEMEncodePrivateKey(privKey ed25519.PrivateKey) (priv []byte, err error) {
	privBytes, err := x509.MarshalPKCS8PrivateKey(privKey)
	if err != nil {
		return nil, err
	}

	pemEncodedPrivKey := pem.EncodeToMemory(&pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: privBytes,
	})

	return pemEncodedPrivKey, nil
}

// PEMEncodeKeyPair encodes an ed25519 keypair using the common PKIX and PKCS8 standards
// with a PEM format
func PEMEncodeKeyPair(pubKey ed25519.PublicKey, privKey ed25519.PrivateKey) (pub []byte, priv []byte, err error) {
	pubKeyBytes, err := PEMEncodePublicKey(pubKey)

	if err != nil {
		return nil, nil, err
	}

	privKeyBytes, err := PEMEncodePrivateKey(privKey)

	if err != nil {
		return nil, nil, err
	}

	return pubKeyBytes, privKeyBytes, nil
}

// PEMDecodePublicKey decodes an ed25519 PEM encoded with PKIX standard
func PEMDecodePublicKey(pub []byte) (ed25519.PublicKey, error) {
	pubBlock, _ := pem.Decode(pub)
	if pubBlock == nil {
		return nil, fmt.Errorf("no pem block found on public key")
	}

	pubKey, err := x509.ParsePKIXPublicKey(pubBlock.Bytes)
	if err != nil {
		return nil, err
	}

	edPubKey, ok := pubKey.(ed25519.PublicKey)
	if !ok {
		return nil, fmt.Errorf("public key is not an ed25519 key")
	}

	return edPubKey, nil
}

// PEMDecodePrivateKey decodes an ed25519 PEM encoded with PKCS8 standard
func PEMDecodePrivateKey(priv []byte) (ed25519.PrivateKey, error) {
	privBlock, _ := pem.Decode(priv)
	if privBlock == nil {
		return nil, fmt.Errorf("no pem block found on private key")
	}

	privKey, err := x509.ParsePKCS8PrivateKey(privBlock.Bytes)
	if err != nil {
		return nil, err
	}

	edPrivKey, ok := privKey.(ed25519.PrivateKey)
	if !ok {
		return nil, fmt.Errorf("private key is not an ed25519 key")
	}

	return edPrivKey, nil
}

// PEMDecodeKeyPair decodes an ed25519 PEM keypair encoded with PKIX and PKCS8 standards
func PEMDecodeKeyPair(pub []byte, priv []byte) (ed25519.PublicKey, ed25519.PrivateKey, error) {
	edPubKey, err := PEMDecodePublicKey(pub)
	if err != nil {
		return nil, nil, err
	}

	edPrivKey, err := PEMDecodePrivateKey(priv)
	if err != nil {
		return nil, nil, err
	}

	return edPubKey, edPrivKey, nil
}

func SignChallenge(
	salt string,
	challenge string,
	password string,
) ([]byte, error) {
	var err error
	var privateKey ed25519.PrivateKey
	h := sha256.New()
	h.Write([]byte(password + salt))
	seed := h.Sum(nil)

	if _, privateKey, err = GenerateKeyPairFromSeed(seed); err != nil {
		return nil, err
	}

	signedChallenge := ed25519.Sign(privateKey, []byte(challenge))

	return signedChallenge, nil
}
