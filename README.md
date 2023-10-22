# go-avial

Go client for AvesTerra. It connects over HGTP, exposes entity operations, and
converts retrieved data into Go types. The root package provides the AVIAL API;
`hgtp/` contains the protocol client and connection pool.

## Install

Requires Go 1.19 or later.

```sh
go get github.com/yanis-fourel/go-avial
```

## Connect and retrieve

You'll need a running AvesTerra server, its certificates, an authorization token,
and the ID of an existing object. Set `AVESTERRA_TOKEN` to your token and replace
the address, certificate directory, and entity ID below with your server's values.

```go
package main

import (
	"fmt"
	"log"
	"os"

	av "github.com/yanis-fourel/go-avial"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	auth, err := av.ParseToken(os.Getenv("AVESTERRA_TOKEN"))
	if err != nil {
		return fmt.Errorf("parse token: %w", err)
	}

	err = av.Initialize(&av.ClientConfig{
		Address:        "127.0.0.1",
		Port:           20057,
		CertificateDir: "/AvesTerra/Certificates/",
	})
	if err != nil {
		return fmt.Errorf("connect: %w", err)
	}
	defer av.Finalize()

	entity := av.NewEntity(0, 0, 123) // Replace with your object's PID, HID, UID.
	model, err := entity.Retrieve(auth)
	if err != nil {
		return fmt.Errorf("retrieve: %w", err)
	}
	fmt.Printf("%+v\n", model)
	return nil
}
```

`Initialize` opens a shared client; call it once and `Finalize` when you're done.
`Retrieve` returns an `AvialModel` decoded from the server's JSON response.

## Tests

Run the model serialization tests without a server:

```sh
go test . -run 'Unmarshal'
```

`go test ./...` also runs integration tests, including the model round-trip test.
These expect a local AvesTerra server, certificates, and the test authorization
configured in `common_test.go` and `hgtp_tests/client_test.go`.
