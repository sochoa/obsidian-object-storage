package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	rootCmd = &cobra.Command{
		Use:   "obsidian",
		Short: "Object storage as a service",
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

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.obsidian.yaml)")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile) // Use config file from the flag.
	} else {
		home, err := homedir.Dir() // Find home directory.
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		viper.AddConfigPath(home) // Search config in home directory with name ".obsidian" (without extension).
		viper.SetConfigName(".obsidian")
	}

	viper.AutomaticEnv() // read in environment variables that match

	if err := viper.ReadInConfig(); err == nil { // If a config file is found, read it in.
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
