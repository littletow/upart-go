package cmd

import (
	"fmt"
	"gart/service"
	"gart/utils"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(qCmd)
}

var qCmd = &cobra.Command{
	Use:   "search",
	Short: "查找文章，最多返回20条记录。",
	Long:  `查找文章， 根据文章的标题和关键字匹配查询，最多返回20条记录。可以输入更多的内容进行精确查找。`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		content := args[0]

		list, err := service.SearchArt(token, content)
		if err != nil {
			fmt.Println("查询发生错误,", err)
		} else {
			n := len(list)
			if n > 0 {
				var (
					cts    string
					uts    string
					ispub  string = "否"
					islock string = "否"
				)

				t := table.NewWriter()
				t.SetOutputMirror(os.Stdout)
				t.AppendHeader(table.Row{"UUID", "题目", "关键字", "是否公开", "是否加锁", "创建时间", "修改时间"})
				for _, v := range list {
					cts = utils.TS2Str(v.Createtime)
					uts = utils.TS2Str(v.Updatetime)
					if v.IsPub == 1 {
						ispub = "是"
					}
					if v.IsLock == 1 {
						islock = "是"
					}
					t.AppendRows([]table.Row{
						{v.Uuid, v.Title, v.Keyword, ispub, islock, cts, uts},
					})
					t.AppendSeparator()
				}

				t.Render()
			} else {
				fmt.Println("未找到记录，尝试更换内容试试。")
			}
		}

	},
}
