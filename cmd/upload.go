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
	rootCmd.AddCommand(upCmd)
}

var upCmd = &cobra.Command{
	Use:   "upload",
	Short: "上传文章，可在豆子碎片小程序中查看。",
	Long: `上传文章，可在豆子碎片小程序中查看。参数依次为题目，关键字，Markdown文件，是否公开，是否加锁。
	参数题目，关键字，Markdown文件必填，是否公开，是否加锁选填，默认为否。
	例如：gart upload 上传示例1 命令行，工具 ./example.md 
	`,
	Args: cobra.RangeArgs(3, 5),
	Run: func(cmd *cobra.Command, args []string) {
		CheckBindAccount()
		var (
			title    string
			keyword  string
			filename string
			ispub    int
			islock   int
		)
		title = args[0]
		keyword = args[1]
		filename = args[2]
		l := len(args)
		switch l {
		case 4:
			ispub = utils.Str2Int(args[3])
		case 5:
			ispub = utils.Str2Int(args[3])
			islock = utils.Str2Int(args[4])
		}
		fmt.Println("上传参数如下：")
		var isPubStr string = "否"
		var isLockStr string = "否"
		if ispub == 1 {
			isPubStr = "是"
		}
		if islock == 1 {
			isLockStr = "是"
		}
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"题目", "关键字", "文件名称", "是否公开", "是否加锁"})
		t.AppendRows([]table.Row{
			{title, keyword, filename, isPubStr, isLockStr},
		})
		t.AppendSeparator()
		t.Render()
		err := service.UploadArt(token, title, keyword, filename, ispub, islock)
		if err != nil {
			fmt.Println("上传发生错误,", err)
		} else {
			fmt.Println("上传成功")
		}
	},
}
