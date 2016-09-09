package cmd

import (
	"github.com/Sirupsen/logrus"

	"github.com/spf13/cobra"
)

var rootArgs = []string{"init", "deploy", "destroy"}

// RootCmd :
var RootCmd = &cobra.Command{
	Use:       "pepper",
	Short:     "Wrapper around salt-cloud",
	Long:      `pepper is a wrapper around salt-cloud that will generate salt-cloud profiles and cloud-init configs (for CoreOS)`,
	ValidArgs: rootArgs,
}

// Execute :
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		logrus.Fatalf("couldn't execute pepper cmd: %v", err)
	}
}
