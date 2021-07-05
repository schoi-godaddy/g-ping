package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "(Short) Print the version number of G-ping",
	Long:  `(Long) Print the version number of G-ping`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("G-ping v0.0.1")
		fmt.Println("args", args)
	},
}
