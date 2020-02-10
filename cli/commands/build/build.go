package build

import (
	"github.com/spf13/cobra"
)

var (
	stack   string
	noCache bool
)

var BuildCommand = &cobra.Command{
	Use:     "build",
	Short:   "Builds a Streamlit ready Docker image from the current folder sources",
	Example: `kubestream build`,
	Run: func(cmd *cobra.Command, args []string) {
		//
	},
}

func init() {
	BuildCommand.Flags().StringVarP(&stack, "stack", "s", "python3-8", "canned environment for the build (optional)")
	BuildCommand.Flags().BoolVar(&noCache, "no-cache", false, "disable Docker cache")
}
