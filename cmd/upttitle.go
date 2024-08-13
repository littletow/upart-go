package cmd

import (
	"fmt"
	"gart/service"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(uptTitleCmd)
}

var uptTitleCmd = &cobra.Command{
	Use:   "title",
	Short: "更新文章标题，参数需要UUID，新的标题。",
	Long:  `更新文章标题，参数需要UUID，新的标题，需要先获取文章的UUID。`,
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		uuid := args[0]
		title := args[1]
		uar := service.UpdateArtReq{
			Uuid:    uuid,
			Title:   title,
			UptType: 1,
		}
		err := service.UpdateArt(token, &uar)
		if err != nil {
			fmt.Println("更新标题发生错误,", err)
		} else {
			fmt.Println("更新成功")
		}
	},
}
