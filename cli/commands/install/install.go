package install

import (
	"fmt"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/relingan/kubestream/cli/util"
	"github.com/spf13/cobra"
)

const (
	defaultNamespace        = "kubestream"
	defaultRegistryImage    = "registry"
	defaultRegistryImageTag = "2"
	defaultDeployment       = "registry"
	defaultService          = "registry"
)

// InstallCommand checks if the cluster is
// running the resources we need to execute
// kubestream deployments.
//
// The necessary workloads are:
// - Docker image registry: This involves a k8s deployment
//   object that uses the https://hub.docker.com/_/registry
//   image, and a k8s service that we will use to expose
//   the registry to kubestream and allow pushing images from
//   the local machine using k8s port forwarding capabilities.
var InstallCommand = &cobra.Command{
	Use:   "install",
	Short: "Bootstrap the resources we need to run Kubestream in your k8s cluster",

	Example: `kubestream install`,
	Run: func(cmd *cobra.Command, args []string) {
		// Get a k8s ClientSet
		cs, err := util.GetClientSet()
		if err != nil {
			fmt.Fprintln(os.Stderr, "error creating k8s client: "+err.Error())
			os.Exit(1)
		}

		spew.Dump(cs)
	},
}
