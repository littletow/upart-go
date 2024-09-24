package cmd

import (
	"fmt"
	"gart/service"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(uptAreaCmd)
}

var uptAreaCmd = &cobra.Command{
	Use:   "city",
	Short: "限制文章为同城访问",
	Long:  `限制文章为同城访问，参数需要UUID，城市名称。`,
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		uuid := args[0]
		city := args[1]

		uar := service.UpdateArtReq{
			Uuid:     uuid,
			AreaFlag: 3,
			AreaCont: city,
			UptType:  6,
		}
		err := service.UpdateArt(token, &uar)
		if err != nil {
			fmt.Println("更新公开开关发生错误,", err)
		} else {
			fmt.Println("更新成功")
		}
	},
}
