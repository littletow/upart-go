package cmd

import (
	"fmt"
	"gart/service"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(rmCmd)
}

var rmCmd = &cobra.Command{
	Use:   "remove",
	Short: "删除文章",
	Long:  `删除文章， 需要先获取文章的UUID。`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		CheckBindAccount()
		uuid := args[0]
		err := service.RemoveArt(token, uuid)
		if err != nil {
			fmt.Println("删除发生错误,", err)
		} else {
			fmt.Println("删除成功")
		}
	},
}
