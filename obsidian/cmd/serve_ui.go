package cmd

import (
	"github.com/sochoa/obsidian/static"
	"github.com/sochoa/obsidian/static/config"
	"github.com/spf13/cobra"
)

var (
	staticConfig config.StaticConfig = config.NewStaticConfig()
	serveUiCmd   *cobra.Command      = &cobra.Command{
		Use:   "serve-ui",
		Short: "Start the object storage ui",
		Run: func(cmd *cobra.Command, args []string) {
			static.Serve(staticConfig)
		},
	}
)

func init() {
	serveUiCmd.PersistentFlags().StringVarP(&staticConfig.Root, "static-root", "", staticConfig.Root, "")
	serveUiCmd.PersistentFlags().StringVarP(&staticConfig.Host, "host", "", staticConfig.Host, "")
	serveUiCmd.PersistentFlags().IntVarP(&staticConfig.Port, "port", "", staticConfig.Port, "")
	rootCmd.AddCommand(serveUiCmd)
}
