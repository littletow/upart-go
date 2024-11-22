package cmd

import (
	"fmt"
	"net/http"

	"github.com/inconshreveable/go-update"
	"github.com/spf13/cobra"
)

const (
	UpURL = "升级URL"
)

var upgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "升级客户端",
	Long:  `升级客户端`,
	// Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if isNewVersion {
			fmt.Println("开始升级")
			err := doUpdate(UpURL)
			if err != nil {
				fmt.Printf("升级失败，%v\n", err)
			} else {
				fmt.Println("升级完毕")
			}

		} else {
			fmt.Println("已是最新版本，无需更新")
		}
	},
}

func init() {
	rootCmd.AddCommand(upgradeCmd)
}

func doUpdate(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	err = update.Apply(resp.Body, update.Options{})
	if err != nil {
		return err
	}
	return nil
}
