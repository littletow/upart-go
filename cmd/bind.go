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
		// 连接Websocket，获取识别码
		// 调用ebtien代码，显示图片，让用户扫码，扫码后会进入mp，然后填写识别码，进行绑定。
		// 匹配成功后，写入到配置文件，并返回结果。
		// 绑定成功后，图片应用退出。结束绑定。
		fmt.Println("实现插入配置文件")
		service.ShowImage()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
