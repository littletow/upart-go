package cmd

import (
	"archive/zip"
	"bytes"
	"crypto"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/inconshreveable/go-update"
	"github.com/spf13/cobra"
)

const (
	UpURL       = "https://gitee.com/littletow/upart-go/releases/download"
	BinFileName = "gart.exe"
	TxtFileName = "sha256.txt"
	ZipFileName = "gart.zip"
)

var upgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "升级客户端",
	Long:  `升级客户端`,
	// Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if isNewVersion {
			fmt.Println("开始升级")
			err := doUpdate()
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

func doUpdate() error {
	// 下载源文件
	url := fmt.Sprintf("%s/v%s/%s", UpURL, gVersion, "gart-win10.zip")
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	// 获取压缩包
	out, err := os.Create(ZipFileName)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}
	chksum, binfile, err := unzip(ZipFileName)
	if err != nil {
		return err
	}
	if len(binfile) == 0 || chksum == "" {
		return errors.New("无效的zip包")
	}
	// 检查
	br := bytes.NewReader(binfile)
	err = updateWithChecksum(br, chksum)
	if err != nil {
		return err
	}
	return nil
}

// 解压缩，然后返回chksum，二进制文件
func unzip(src string) (string, []byte, error) {
	var (
		chksum  string
		binfile []byte
	)
	r, err := zip.OpenReader(src)
	if err != nil {
		return chksum, binfile, err
	}
	defer r.Close()

	for _, f := range r.File {
		fileName := f.Name
		fmt.Printf("Contents of %s:\n", fileName)
		if fileName == BinFileName || fileName == TxtFileName {
			if fileName == BinFileName {
				rc, err := f.Open()
				if err != nil {
					log.Println(fileName, "open error,", err)
				} else {
					ff, err := io.ReadAll(rc)
					if err != nil {
						log.Println(fileName, "read error,", err)
					} else {
						binfile = ff
					}
				}
				rc.Close()
			}

			if fileName == TxtFileName {
				rc, err := f.Open()
				if err != nil {
					log.Println(fileName, "open error,", err)
				} else {
					ff, err := io.ReadAll(rc)
					if err != nil {
						log.Println(fileName, "read error,", err)
					} else {
						chksum = string(ff)
						chksum = strings.TrimSpace(chksum)
					}
				}
				rc.Close()
			}
		}
	}
	return chksum, binfile, nil
}

func updateWithChecksum(binary io.Reader, hexChecksum string) error {
	checksum, err := hex.DecodeString(hexChecksum)
	if err != nil {
		return err
	}
	err = update.Apply(binary, update.Options{
		Hash:     crypto.SHA256, // this is the default, you don't need to specify it
		Checksum: checksum,
	})

	return err
}
