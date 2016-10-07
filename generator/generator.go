package generator

import (
	"bytes"
	"crypto/rand"
	"errors"
	"math/big"
	"strings"
)

/**
 * Set the contant letters/numbers/specials
 */
const (
	lowerCaseCharacters = "abcdefghijklmnopqrstuvwxyz"
	numbers             = "0123456789"
	specialCharacters   = "$%&*@^"
)

func NewPassword(length int, uppercase bool, lowercase bool, number bool, special bool, ) (password string, error error) {
	var passwordBuffer bytes.Buffer

	// Use full character set if non specified
	if false == lowercase && false == uppercase && false == number && false == special {
		lowercase = true
		uppercase = true
		number    = true
	}

	// Add lower case to the character set
	if lowercase {
		passwordBuffer.WriteString(lowerCaseCharacters)
	}

	// Add upper case to the character set
	if uppercase {
		passwordBuffer.WriteString(strings.ToUpper(lowerCaseCharacters))
	}

	// Add numbers to the character set
	if number {
		passwordBuffer.WriteString(numbers)
	}

	// Add special chars to the character set
	if special {
		passwordBuffer.WriteString(specialCharacters)
	}

	if passwordBuffer.Len() <= 0 {
		error = errors.New("Password must be at least one character long.")
		return
	}

	// Generate the password
	chars := passwordBuffer.String()
	limit := big.NewInt(int64(len(passwordBuffer.String())))
	passwordChars := make([]byte, length)
	for i := 0; i < length; i++ {
		randNum, error := rand.Int(rand.Reader, limit)
		if error != nil {
			return "", error
		}
		passwordChars[i] = chars[randNum.Int64()]
	}

	password = string(passwordChars)

	return
}