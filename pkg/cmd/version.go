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
	Short: "Print the version number of this CLI",
	Long:  `All software has versions. This is Nitric's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Nitric Helper CLI v2.0")
	},
}