package cmd

import (
	"github.com/sochoa/obsidian/crud/config"
	"github.com/spf13/cobra"
)

var (
	serveConfig config.ObjectStorageConfig
	serveCmd    *cobra.Command = &cobra.Command{
		Use:   "serve",
		Short: "Start the object storage service",
		Run: func(cmd *cobra.Command, args []string) {
			crud.Serve(serveConfig)
		},
	}
)

func init() {
	serveConfig = crud.NewObjectStorageConfig()
	serveCmd.PersistentFlags().StringVarP(&serveConfig.StorageRoot, "storage-root", "", serveConfig.StorageRoot, "Where objects are stored on the local filesystem")
	serveCmd.PersistentFlags().StringVarP(&serveConfig.Host, "host", "", serveConfig.Host, "")
	serveCmd.PersistentFlags().IntVarP(&serveConfig.Port, "port", "", serveConfig.Port, "")
	rootCmd.AddCommand(serveCmd)
}
