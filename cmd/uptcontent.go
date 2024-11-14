package cmd

import (
	"fmt"
	"gart/service"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(uptContentCmd)
}

var uptContentCmd = &cobra.Command{
	Use:   "content",
	Short: "更新文章内容，参数需要UUID，新的文件内容。",
	Long:  `更新文章内容，参数需要UUID，新的文件内容，需要先获取文章的UUID。`,
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		CheckBindAccount()
		uuid := args[0]
		filename := args[1]
		content, err := service.GetFileContent(filename)
		if err != nil {
			fmt.Println("读取文件内容发生错误,", err)
			return
		}
		uar := service.UpdateArtReq{
			Uuid:    uuid,
			Content: string(content),
			UptType: 3,
		}
		err = service.UpdateArt(token, &uar)
		if err != nil {
			fmt.Println("更新内容发生错误,", err)
		} else {
			fmt.Println("更新成功")
		}
	},
}
