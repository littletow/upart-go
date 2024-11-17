package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

// 每次新增功能都增一个版本，bug修复增一个小版本。
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "获取版本号",
	Long:  `获取版本号`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("当前版本号：%s\n", "v3.0.0")
		fmt.Println("版本 v3.0.0 特性如下：")
		fmt.Println("1. 支持检测是否初始化配置，自动创建配置文件，绑定豆子碎片识别码")
		fmt.Println("2. 支持查看 Markdown 用法")
		fmt.Println("3. 支持豆子点数不足时提示看广告")
	},
}
