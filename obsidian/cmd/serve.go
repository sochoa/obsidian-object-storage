package cmd

import (
	"github.com/sochoa/obsidian/crud"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the object storage service",
	Run: func(cmd *cobra.Command, args []string) {
		crud.Serve()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
