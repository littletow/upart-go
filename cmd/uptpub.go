package cmd

import (
	"fmt"
	"gart/service"
	"gart/utils"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(uptPublicCmd)
}

var uptPublicCmd = &cobra.Command{
	Use:   "public",
	Short: "更新文章是否公开，参数需要UUID，新的公开开关，只允许0或1,1为公开。",
	Long:  `更新文章是否公开，参数需要UUID，新的公开开关，只允许0或1,1为公开，需要先获取文章的UUID。`,
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		CheckBindAccount()
		uuid := args[0]
		ispub := utils.Str2Int(args[1])
		uar := service.UpdateArtReq{
			Uuid:    uuid,
			IsPub:   ispub,
			UptType: 4,
		}
		err := service.UpdateArt(token, &uar)
		if err != nil {
			fmt.Println("更新公开开关发生错误,", err)
		} else {
			fmt.Println("更新成功")
		}
	},
}
