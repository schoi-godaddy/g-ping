package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	Verbose bool
	rootCmd = &cobra.Command{
		Use:   "gping",
		Short: "G-ping is a small graphical ping application in Go.",
		Long: `Hello! Welcome to my first G-ping CLI tool!
	This was created with Golang & Cobra.
	First time implementing something like this, so enjoy!`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("hello?")
			fmt.Println("args", args)
			fmt.Println("Verbose", Verbose)
		},
	}
)

func init() {
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "Output in Verbose mode")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
