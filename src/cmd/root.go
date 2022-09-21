package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/brandtkeller/oscal-builder/src/cmd/create"
)

var rootCmd = &cobra.Command{
	Use:   "oscal-builder",
	Short: "oscal-builder",
	Long:  `oscal-builder`,
}

func Execute() {

	commands := []*cobra.Command{
		create.Command(),
	}

	rootCmd.AddCommand(commands...)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
