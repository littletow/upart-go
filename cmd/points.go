package cmd

import (
	"fmt"
	"gart/service"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(pointsCmd)
}

// 每次新增功能都增一个版本，bug修复增一个小版本。
var pointsCmd = &cobra.Command{
	Use:   "bean",
	Short: "获取豆子点数数量",
	Long:  `获取拥有的豆子点数数量`,
	Run: func(cmd *cobra.Command, args []string) {
		CheckBindAccount()
		points, err := service.GetPoints(token)
		if err != nil {
			fmt.Println("获取豆子点数，", err)
			os.Exit(1)
		}
		fmt.Printf("您拥有豆子点数 %d 个\n", points)

	},
}
