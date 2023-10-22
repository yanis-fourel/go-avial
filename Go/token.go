package av

import "github.com/yanis-fourel/go-avial/Go/hgtp"

type Token = hgtp.Token

var NullToken = hgtp.NullToken
var ViewToken = hgtp.ViewToken
var NoToken = hgtp.NoToken
var TokenObfuscated = hgtp.TokenObfuscated

// Decodes str into a UUID or returns an error.
// Both the standard UUID forms of xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx and
// urn:uuid:xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx are decoded as well as the
// Microsoft encoding {xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx} and the raw hex
// encoding: xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx.
func ParseToken(str string) (Token, error) {
	return hgtp.ParseToken(str)
}

// Like `ParseToken` but panics if the string cannot be parsed.
// It simplifies safe initialization of global variables holding compiled UUIDs.
func MustParseToken(str string) Token {
	return hgtp.MustParseToken(str)
}

// Creates a new Authorization from a byte slice.
// Returns an error if the slice does not have a length of 16.
// The bytes are copied from the slice.
func TokenFromBytes(b []byte) (Token, error) {
	return hgtp.TokenFromBytes(b)
}

// Creates a new random Token or panics
func RandomToken() Token {
	return hgtp.RandomToken()
}
