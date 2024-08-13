package cmd

import (
	"fmt"
	"gart/service"
	"gart/utils"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(uptLockCmd)
}

var uptLockCmd = &cobra.Command{
	Use:   "lock",
	Short: "更新文章是否加锁，参数需要UUID，新的加锁开关，只允许0或1,1为加锁。",
	Long:  `更新文章是否加锁，参数需要UUID，新的加锁开关，只允许0或1,1为加锁，需要先获取文章的UUID。`,
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		uuid := args[0]
		islock := utils.Str2Int(args[1])
		uar := service.UpdateArtReq{
			Uuid:    uuid,
			IsLock:  islock,
			UptType: 5,
		}
		err := service.UpdateArt(token, &uar)
		if err != nil {
			fmt.Println("更新加锁开关发生错误,", err)
		} else {
			fmt.Println("更新成功")
		}
	},
}
