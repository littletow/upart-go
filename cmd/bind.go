package cmd

import (
	"fmt"
	"net"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(initCmd)
}

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
	},
}

// 获取本机MAC地址
func GetMacAddr() string {
	ints, err := net.Interfaces()
	if err != nil {
		panic("无法获取本机网络，" + err.Error())
	}

	i0 := ints[0]
	mac := i0.HardwareAddr.String()
	return mac
}

// 连接Websocket
func ConnWs() {
	// 获取识别码，并接收绑定结果

}

// 显示图片
func ShowImage() {
	// 使用ebtien显示

}
