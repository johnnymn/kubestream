package build

import (
	"github.com/spf13/cobra"
)

var dockerfile, requirements, environment string

var BuildCommand = &cobra.Command{
	Use:     "build <file>",
	Short:   "Builds a Docker image that can serve the given Streamlit application",
	Example: `kubestream build app.py`,
	Run: func(cmd *cobra.Command, args []string) {
		//
	},
}

func init() {
	BuildCommand.Flags().StringVarP(&dockerfile, "dockerfile", "d", "", "path to Dockerfile (optional)")
	BuildCommand.Flags().StringVarP(&environment, "environment", "e", "python3-8", "bundled environment for the build (optional)")
	BuildCommand.Flags().StringVarP(&requirements, "requirements", "r", "", "path to requirements.txt file (optional)")
}
