package cmd

import (
	"fmt"
	"gart/service"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(uptKeywordCmd)
}

var uptKeywordCmd = &cobra.Command{
	Use:   "keyword",
	Short: "更新文章关键字，参数需要UUID，新的关键字。",
	Long:  `更新文章关键字，参数需要UUID，新的关键字，需要先获取文章的UUID。`,
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		uuid := args[0]
		keyword := args[1]
		uar := service.UpdateArtReq{
			Uuid:    uuid,
			Keyword: keyword,
			UptType: 2,
		}
		err := service.UpdateArt(token, &uar)
		if err != nil {
			fmt.Println("更新关键字发生错误,", err)
		} else {
			fmt.Println("更新成功")
		}
	},
}
