package github

import (
	"bytes"
	"log"

	"github.com/cli/go-gh/v2"
)

func UploadSshPublicKey(title, path string) (bytes.Buffer, error) {
	_, queryResponse, err := gh.Exec("ssh-key", "add", path, "--title", title, "--type", "authentication")
	if err != nil {
		log.Fatal(err)
	}
	return queryResponse, nil
}
