package cmd

import (
	"errors"
	"fmt"
	"gart/service"
	"os"
	"path"
	"time"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const intro = `gart是一个上传豆子碎片文章和管理文章的一个命令行工具。
它使用 Golang 实现。

已经实现的功能（命令）如下：
1. upload 上传文章
2. remove 文章删除
3. search 根据标题或者关键字查找文章
4. title 更新文章标题
5. keyword 更新文章关键字
6. content 更新文章内容
7. public 将文章公开
8. lock 将文章加锁
9. forcepub 将文章强制公开
10. init 绑定账号，初始化配置文件
11. version 打印版本号
12. markdown 获取Markdown常用语法教程
13. miniapp 获取豆子碎片小程序码
14. area 获取有效省份和城市
15. city 限制文章为同城访问

gart使用语法：gart 命令
或使用gart --help获取帮助`

var (
	cfgFile  string
	token    string
	isEnable bool
)

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/upart.toml)")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		// 检查文件是否存在？
		fileName := path.Join(home, "upart.toml")
		_, err = os.Stat(fileName)
		if err != nil {
			// 第一次创建
			_, err = os.Create(fileName)
			if err != nil {
				fmt.Println("无法创建配置文件", err)
				os.Exit(1)
			}
			viper.AddConfigPath(home)
			viper.SetConfigName("upart")
			viper.SetConfigType("toml")
			// 设置默认值
			viper.SetDefault("expire_at", 0)
			viper.SetDefault("token", "")
			viper.SetDefault("icode", "")
			viper.SetDefault("isecret", "")
			viper.SetDefault("is_enable", false)
			err = viper.WriteConfig()
			if err != nil {
				fmt.Println("写入配置文件错误,", err)
				os.Exit(1)
			}
		} else {
			viper.AddConfigPath(home)
			// viper.AddConfigPath(".")
			viper.SetConfigName("upart")
			viper.SetConfigType("toml")
		}
	}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Cant't read config:", err)
		os.Exit(1)
	}
	// 每次启动都调用token
	isEnable = viper.GetBool("is_enable")
	if isEnable {
		err := GetToken()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		// points, err = service.GetPoints(token)
		// if err != nil {
		// 	fmt.Println(err)
		// 	os.Exit(1)
		// }
	}
}

func GetToken() error {
	var err error
	icode := viper.GetString("icode")
	isecret := viper.GetString("isecret")
	expireAt := viper.GetInt64("expire_at")
	now := time.Now().Unix()
	if icode == "" && isecret == "" {
		return errors.New("警告：还未绑定账户，请使用 `gart init` 命令初始化配置。")
	}
	if now > expireAt {
		token, err = service.GetToken(icode, isecret)
		if err != nil {
			fmt.Println("获取token错误,", err)
			return err
		}
		viper.Set("expire_at", now+7000)
		viper.Set("token", token)
		err = viper.WriteConfig()
		if err != nil {
			fmt.Println("写入配置文件错误,", err)
			return err
		}
	} else {
		token = viper.GetString("token")
	}
	return nil
}

var rootCmd = &cobra.Command{
	Use:   "gart",
	Short: "gart 是文章管理命令行工具。",
	Long:  `gart 是文章管理命令行工具，主要用来管理豆子碎片小程序中的文章。`,
	Run: func(cmd *cobra.Command, args []string) {
		// CheckBindAccount()
		fmt.Println(intro)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
