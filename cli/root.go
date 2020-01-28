package cmd

import (
	"fmt"
	"os"

	"github.com/relingan/kubestream/cli/commands/build"
	"github.com/relingan/kubestream/cli/commands/install"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "kubestream <command>",
	Short: "CLI tool to deploy Streamlit apps to Kubernetes",
	Long: `
kubestream takes care of packaging your Streamlit app into
a Docker container and uses the Kubernetes API to deploy
it to your cluster without you having to worry about managing
the underlaying resources.
`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(build.BuildCommand)
	rootCmd.AddCommand(install.InstallCommand)
}
