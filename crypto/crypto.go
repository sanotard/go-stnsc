// crypto package provide linux user password verification method
package crypto

import (
	"errors"
	"strings"

	"github.com/kless/osutil/user/crypt"
	"github.com/kless/osutil/user/crypt/apr1_crypt"
	"github.com/kless/osutil/user/crypt/md5_crypt"
	"github.com/kless/osutil/user/crypt/sha256_crypt"
	"github.com/kless/osutil/user/crypt/sha512_crypt"
)

var (
	// Password verification failed
	ErrVerificationFailed = errors.New("Password verification failed")
	// Password format is inccorrenct
	ErrIncorrectFormat = errors.New("Incorrect password format")
)

// Verify hashedPass and rawPass.
// hashedPass is linux encrypted passowrd like /etc/shadow.
// rawPass is raw string pasword cast to binary.
//
// e.g.
//  hashedPass : $6$RNqhn2ttIfMcRj4r$Ddnbckw1T1xUkguDWvSsb3GZseoeahRbr27vKbYV9opja2SKWi6y.67YI0yXz8HremKCpJwwFEOqed6Eic9.0.
//  rawPass : password123
//
// Support encryption algorithm - SHA-512, HSA-256, MD5 and APR1
func Verify(hashedPass string, rawPass []byte) error {
	crypter, err := crypter(hashedPass)
	if err != nil {
		return err
	}

	err = crypter.Verify(hashedPass, rawPass)
	if err != nil {
		return ErrVerificationFailed
	}

	return nil

}

// Return crypter by checking hashedPass.
func crypter(hashedPass string) (crypt.Crypter, error) {
	var c crypt.Crypter
	switch {
	case strings.HasPrefix(hashedPass, sha512_crypt.MagicPrefix):
		c = sha512_crypt.New()
	case strings.HasPrefix(hashedPass, sha256_crypt.MagicPrefix):
		c = sha256_crypt.New()
	case strings.HasPrefix(hashedPass, md5_crypt.MagicPrefix):
		c = md5_crypt.New()
	case strings.HasPrefix(hashedPass, apr1_crypt.MagicPrefix):
		c = apr1_crypt.New()
	default:
		return nil, ErrIncorrectFormat
	}
	return c, nil
}
