package main

import "fmt"

type PasswordProtector struct {
	user          string
	passwordName  string
	hashAlgorithm HashAlgorithm
}

type HashAlgorithm interface {
	Hash(passwordProtector *PasswordProtector)
}

func NewPasswordProtector(user string, passwordName string, HashAlgorithm HashAlgorithm) *PasswordProtector {
	return &PasswordProtector{
		user:          user,
		passwordName:  passwordName,
		hashAlgorithm: HashAlgorithm,
	}
}

func (passwordProtector *PasswordProtector) SetHashAlgorithm(hashAlgorithm HashAlgorithm) {
	passwordProtector.hashAlgorithm = hashAlgorithm
}

func (passwoedProtector *PasswordProtector) Hash() {
	passwoedProtector.hashAlgorithm.Hash(passwoedProtector)
}

type SHA struct{}

func (SHA) Hash(passwordProtector *PasswordProtector) {
	fmt.Printf("Hashing using SHA for %s\n", passwordProtector.passwordName)
}

type MD5 struct{}

func (MD5) Hash(passwordProtector *PasswordProtector) {
	fmt.Printf("Hashing using MD5 for %s\n", passwordProtector.passwordName)
}

func main() {
	sha := &SHA{}
	md5 := &MD5{}

	passwordProtector := NewPasswordProtector("pelli", "gmail password", sha)
	passwordProtector.Hash()
	passwordProtector.SetHashAlgorithm(md5)
	passwordProtector.Hash()
}
