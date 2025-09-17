
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of transcat",
	Long: "All software has versions. This is transcat",
	Run: func(cmd *cobra.Command, args []string) {
    // ここに処理を書いていく
		fmt.Println("version 0.1.0")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}