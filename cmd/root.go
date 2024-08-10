package cmd

import (
	"fmt"
	"gart/service"
	"gart/utils"
	"os"
	"strconv"
	"time"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	token   string
)

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/conf.toml)")
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
		viper.AddConfigPath(home)
		viper.AddConfigPath(".")
		viper.SetConfigName("conf")
		viper.SetConfigType("toml")
	}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Cant't read config:", err)
		os.Exit(1)
	}
	// 每次启动都调用token
	getToken()

}

func getToken() {
	var err error
	icode := viper.GetString("icode")
	isecret := viper.GetString("isecret")
	expire := viper.GetString("expire_at")
	expireAt := utils.Str2Int64(expire)
	now := time.Now().Unix()
	if now > expireAt {
		token, err = service.GetToken(icode, isecret)
		if err != nil {
			fmt.Println("获取token错误,", err)
			os.Exit(1)
		}
		expireAtStr := strconv.FormatInt(now+7000, 10)
		viper.Set("expire_at", expireAtStr)
		viper.Set("token", token)
		err = viper.WriteConfig()
		if err != nil {
			fmt.Println("写入配置文件错误,", err)
			os.Exit(1)
		}
	} else {
		token = viper.GetString("token")
	}
}

var rootCmd = &cobra.Command{
	Use:   "gart",
	Short: "gart 是文章管理命令行工具。",
	Long:  `gart 是文章管理命令行工具，主要用来管理豆子碎片小程序中的文章。`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("该命令行工具用来管理豆子碎片小程序的文章。")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
