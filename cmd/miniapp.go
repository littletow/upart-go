package cmd

import (
	"fmt"
	"gart/service"

	"github.com/spf13/cobra"
)

var mpCmd = &cobra.Command{
	Use:   "miniapp",
	Short: "获取豆子碎片小程序码",
	Long:  `获取豆子碎片小程序码`,
	// Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("可扫码查看效果")
		service.ShowMpCode()
	},
}

func init() {
	rootCmd.AddCommand(mpCmd)
}
