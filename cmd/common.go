package cmd

import (
	"fmt"
	"gart/service"
	"os"
)

// 检查是否激活账号
func CheckBindAccount() {
	if !isEnable {
		fmt.Println("还未绑定账号，请使用 `gart init` 命令初始化配置。")
		os.Exit(1)
	}
}

func CheckPoints(pt int) {
	points, _ := service.GetPoints(token)
	if points < pt {
		service.RunTui()
	}
}
