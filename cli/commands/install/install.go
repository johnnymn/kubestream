package install

import (
	"fmt"
	"os"

	"github.com/relingan/kubestream/pkg/k8s"
	"github.com/relingan/kubestream/stacks"
	"github.com/spf13/cobra"
)

const (
	v1Stack = "v1"
)

var kubeconfig string

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
	Use:     "install",
	Short:   "Bootstrap the resources we need to run Kubestream in your k8s cluster",
	Example: `kubestream install`,
	Run: func(cmd *cobra.Command, args []string) {
		// Get a k8s clients
		dynClient, err := k8s.GetDynamicClient(kubeconfig)
		if err != nil {
			fmt.Fprintln(os.Stderr, "error creating k8s dynamic client: "+err.Error())
			os.Exit(1)
		}

		clientSet, err := k8s.GetClientSet(kubeconfig)
		if err != nil {
			fmt.Fprintln(os.Stderr, "error creating k8s ClientSet: "+err.Error())
			os.Exit(1)
		}

		// Get the manifests for the stack.
		stack, err := stacks.GetManifests(v1Stack)
		if err != nil {
			fmt.Fprintln(os.Stderr, "error rendering stack: "+err.Error())
			os.Exit(1)
		}

		err = k8s.DeployStack(stack, *dynClient, clientSet)
		if err != nil {
			fmt.Fprintln(os.Stderr, "error deploying stack: "+err.Error())
			os.Exit(1)
		}

		fmt.Print("\nKubestream stack deployed successfully!\n")
	},
}

func init() {
	InstallCommand.Flags().StringVarP(&kubeconfig, "kubeconfig", "k", "", "path to the kubeconfig file (optional)")
}
