package datastore

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/subtle"
	"fmt"
	"io"
	"strings"

	"github.com/ofonimefrancis/pixels/pkg/commons/must"
	"golang.org/x/crypto/scrypt"
)

// Password constants
const (
	SaltLen        = 32
	HashLen        = 64
	MinPasswordLen = 8
)

type WrongPasswordError string

func (self WrongPasswordError) Error() string {
	return string(self)
}

type PasswordHash struct {
	Hash []byte `json:"hash"`
	Salt []byte `json:"salt"`
}

func NewPasswordHash(password string) (*PasswordHash, error) {
	salt := generateSalt()
	hash, err := createPasswordHash(password, salt)
	if err != nil {
		return nil, err
	}

	return &PasswordHash{Hash: hash, Salt: salt}, nil
}

func (self *PasswordHash) IsEqualTo(password string) bool {
	return VerifyPassword(password, self.Hash, self.Salt)
}

// Generate a random salt of suitable length
func generateSalt() []byte {
	salt := make([]byte, SaltLen)
	must.DoF(func() error {
		_, err := rand.Read(salt)
		return err
	})
	return salt
}

// Create a hash of a password and salt
func createPasswordHash(password string, salt []byte) ([]byte, error) {
	password = strings.TrimSpace(password)

	if len(password) < MinPasswordLen {
		return nil, fmt.Errorf("The password is too short")
	}

	hash, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, HashLen)

	if err != nil {
		return nil, err
	}

	return hash, nil
}

// VerifyPassword checks that a password matches a stored hash and salt
func VerifyPassword(password string, hash []byte, salt []byte) bool {
	newhash, err := createPasswordHash(password, salt)
	if err != nil {
		return false
	}
	return subtle.ConstantTimeCompare(newhash, hash) == 1
}

func IsPasswordMatch(password, userpass string) bool {
	ePassword := EncryptPassword(password)

	if ePassword != userpass {
		return false
	}
	return true
}

func EncryptPassword(password string) string {
	tPass := md5.New()
	io.WriteString(tPass, password)
	ePassword := fmt.Sprintf("%x", tPass.Sum(nil))

	return ePassword
}
