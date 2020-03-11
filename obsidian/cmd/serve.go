package cmd

import (
	"github.com/sochoa/obsidian/crud"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var (
	serveCmd *cobra.Command = &cobra.Command{
		Use:   "serve",
		Short: "Start the object storage service",
		Run: func(cmd *cobra.Command, args []string) {
			crud.Serve(GetPwd())
		},
	}
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

func GetPwd() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
		return "/tmp"
	}
	return dir
}
