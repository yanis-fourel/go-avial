package hgtp

import "github.com/google/uuid"

type Token = uuid.UUID

var NullToken = MustParseToken("00000000-0000-0000-0000-000000000000")
var ViewToken = MustParseToken("00000000-0000-0000-0000-000000000001")
var NoToken = MustParseToken("ffffffff-ffff-ffff-ffff-ffffffffffff")
var TokenObfuscated = "********-****-****-****-************"

// Decodes str into a UUID or returns an error.
// Both the standard UUID forms of xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx and
// urn:uuid:xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx are decoded as well as the
// Microsoft encoding {xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx} and the raw hex
// encoding: xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx.
func ParseToken(str string) (Token, error) {
	return uuid.Parse(str)
}

// Like `ParseToken` but panics if the string cannot be parsed.
// It simplifies safe initialization of global variables holding compiled UUIDs.
func MustParseToken(str string) Token {
	return uuid.MustParse(str)
}

// Creates a new Authorization from a byte slice.
// Returns an error if the slice does not have a length of 16.
// The bytes are copied from the slice.
func TokenFromBytes(b []byte) (Token, error) {
	return uuid.FromBytes(b)
}

// Creates a new random Token or panics
func RandomToken() Token {
	return uuid.New()
}
