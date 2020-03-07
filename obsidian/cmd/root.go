package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	binary  = "obsidian"
	version = "development"
	rootCmd = &cobra.Command{
		Use:   binary,
		Short: fmt.Sprintf("%s is an object store", binary),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Do stuff")
		},
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
