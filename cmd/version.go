package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

// 每次新增功能都增一个版本，bug修复增一个小版本。
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "获取版本号",
	Long:  `获取版本号`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("v2.1.0")
	},
}
