package cmd

import (
	"fmt"
	"gart/service"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(forcePublicCmd)
}

var forcePublicCmd = &cobra.Command{
	Use:   "forcepub",
	Short: "强制更新文章公开",
	Long:  `不用通过后台审核，强制更新文章公开，参数需要UUID，消耗10个豆子点数。`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		CheckBindAccount()
		CheckPoints(10)
		uuid := args[0]
		uar := service.UpdateArtReq{
			Uuid:    uuid,
			IsPub:   1,
			UptType: 7,
		}
		err := service.UpdateArt(token, &uar)
		if err != nil {
			fmt.Println("强制公开文章发生错误,", err)
		} else {
			fmt.Println("操作成功")
		}
	},
}
