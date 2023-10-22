package av

import "testing"

func setup(t *testing.T) Token {
	err := Initialize(&ClientConfig{
		Address: "localhost",
	})
	if err != nil {
		t.Fatal(err)
	}

	token := RandomToken()
	rootAuth := MustParseToken("636d384c-54d5-4bab-99b0-9b29236cff59")

	err = RedirectToken(rootAuth, token, rootAuth)
	if err != nil {
		Finalize()
		t.Fatal(err)
	}
	return token
}

func teardown(token Token) {
	err1 := StopRedirectToken(token, MustParseToken("636d384c-54d5-4bab-99b0-9b29236cff59"))
	err2 := Finalize()

	if err1 != nil {
		panic(err1)
	}
	if err2 != nil {
		panic(err2)
	}
}
