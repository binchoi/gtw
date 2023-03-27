/*
Copyright © 2023 BIN CHOI binchoi.kr@gmail.com
*/
package cmd

import (
	"github.com/binchoi/gtw/watcher"
	"os"

	"github.com/spf13/cobra"
)

var (
	// Flag values
	watchPath string

	// rootCmd represents the base command when called without any subcommands
	rootCmd = &cobra.Command{
		Use:   "gtw",
		Short: "A cli tool for automatically running go tests",
		Long: `
	gtw is a CLI library for Go that supercharges test-driven development.
	It is a zero-config cli tool that runs go test automatically when a file 
	in your project is changed.
	`,
		Run: func(cmd *cobra.Command, args []string) {
			watcher.StartWatcher(watchPath)
		},
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVarP(&watchPath, "path", "p", "./...", "Test directory path")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
