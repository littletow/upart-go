package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

// 当前版本为1.2.0，每次新增功能都增一个版本，bug修复增一个小版本。
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "gart version",
	Long:  `gart version`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("v1.2.0")
	},
}