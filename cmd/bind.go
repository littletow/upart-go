package cmd

import (
	"fmt"
	"gart/service"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "绑定账号，初始化配置文件",
	Long:  `绑定账号，初始化配置文件`,
	// Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if !isEnable {
			service.ShowBindCode()
		} else {
			fmt.Println("已绑定账户")
		}

	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
