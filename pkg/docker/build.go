package docker

import (
	"io/ioutil"
	"os"

	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
)

const (
	tmpDirPrefix = "kubestream"
)

type BuildOptions struct{}

func Build(options BuildOptions) error {
	// We need to create a tmp folder to
	// place the Docker build context.
	dir, err := ioutil.TempDir("", tmpDirPrefix)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	defer os.RemoveAll(dir)

	// Copy the Docker stack default files
	// to our tmp folder.

	// Build the image
	return nil
}

func init() {
	log.SetHandler(cli.Default)
}
