package cmd

import (
	"fmt"
	"gart/service"
	"gart/utils"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(validAreaCmd)
}

var validAreaCmd = &cobra.Command{
	Use:   "area",
	Short: "获取有效省份及城市，1为查询省份，2为查询城市，需要提供省份名称",
	Long:  `获取有效省份及城市，1为查询省份，2为查询城市，需要提供省份名称，用于文章同城限制。`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		p1 := utils.Str2Int(args[0])
		if p1 < 1 || p1 > 2 {
			fmt.Println("只允许为1或2")
			return
		}

		if p1 == 2 {
			n := len(args)
			if n < 2 {
				fmt.Println("请提供省份名称")
				return
			}
			p2 := args[1]
			if p2 == "" {
				fmt.Println("请提供省份名称")
				return
			}

			str, err := service.GetValidCity(token, p2)
			if err != nil {
				fmt.Println("查询有效城市发生错误,", err)
			} else {
				fmt.Println("有效城市：", str)
			}
		} else {
			str, err := service.GetValidProvince(token)
			if err != nil {
				fmt.Println("查询有效省份发生错误,", err)
			} else {
				fmt.Println("有效省份：", str)
			}
		}
	},
}
