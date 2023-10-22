package main

import (
	"github.com/yanis-fourel/go-avial"
	"log"
	"os"
)

const CERTIFICATE_PATH = "/AvesTerra/Certificates/"
const KEYLOG_FILE = "/tmp/avesterra-keylog"
const ADDRESS = "127.0.0.1"
const PORT = 20057

func main() {
	log.SetOutput(os.Stdout)
	os.Exit(safemain())
}

func safemain() int {
	keyLogWriter, err := os.Create(KEYLOG_FILE)
	if err != nil {
		log.Print("Error: ", err)
		return 1
	}

	config := av.ClientConfig{
		Address:        ADDRESS,
		Port:           PORT,
		CertificateDir: CERTIFICATE_PATH,
		KeyLogWriter:   keyLogWriter,
	}

	err = av.Initialize(&config)
	if err != nil {
		log.Print("Error: ", err)
		return 1
	}
	defer av.Finalize()

	rootAuth := av.MustParseToken("636d384c-54d5-4bab-99b0-9b29236cff59")

	e, err := av.CreateFolder(av.ArgsCreateFolder{
		Authorization: rootAuth,
	})
	if err != nil {
		log.Print("Error: ", err)
		return 1
	}

	err = e.DeleteFolder(rootAuth)
	if err != nil {
		log.Print("Error: ", err)
		return 1
	}

	return 0
}
