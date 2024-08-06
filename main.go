package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"gopkg.in/ini.v1"
)

type RespData struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data string `json:"data"`
}

type RespMsg struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type ArtReq struct {
	Title   string `json:"title"`
	Keyword string `json:"keyword"`
	Content string `json:"content"`
	IsLock  int    `json:"islock"`
	IsPub   int    `json:"ispub"`
}

const (
	OPEN_URL = "https://open.91demo.top/api/open"
)

var (
	title    string
	keyword  string
	filename string
	ispub    int // 1 pub
	islock   int // 1 lock

)

func init() {
	flag.StringVar(&title, "b", "", "文章题目")
	flag.StringVar(&keyword, "k", "", "文章关键字")
	flag.StringVar(&filename, "f", "", "MD文件")
	flag.IntVar(&islock, "l", 0, "是否加锁")
	flag.IntVar(&ispub, "p", 0, "是否开放")
}

func printHelp() {
	fmt.Println("介绍")
	fmt.Println("./upart.exe -b title -k keyword -f filename -l islock -p ispub")

}

func main() {
	flag.Parse()
	var err error

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println("获取系统目录错误,", err)
		os.Exit(1)
	}
	confFile := filepath.Join(dir, "conf.ini")
	cfg, err := ini.Load(confFile)
	if err != nil {
		fmt.Println("读取配置文件错误,", err)
		os.Exit(1)
	}
	icode := cfg.Section("User").Key("icode").String()
	isecret := cfg.Section("User").Key("isecret").String()
	token := cfg.Section("User").Key("token").String()
	expireAt := cfg.Section("User").Key("expire_at").MustInt64()
	// 刷新token
	now := time.Now().Unix()
	if now > expireAt {
		token, err = getToken(icode, isecret)
		if err != nil {
			fmt.Println("获取token错误,", err)
			os.Exit(1)
		}
		expireAtStr := strconv.FormatInt(now+7000, 10)
		cfg.Section("User").Key("token").SetValue(token)
		cfg.Section("User").Key("expire_at").SetValue(expireAtStr)
		cfg.SaveTo(confFile)
	}
	if token == "" || title == "" || filename == "" {
		printHelp()
		fmt.Println("题目和Markdown文件名不能为空")
		os.Exit(1)
	}
	// 上传文章
	err = uploadArt(token, title, keyword, filename, ispub, islock)
	if err != nil {
		fmt.Println("上传文章错误,", err)
		os.Exit(1)
	}

	fmt.Println("成功上传")
}

func getToken(icode string, isecret string) (string, error) {
	url := fmt.Sprintf("%s/getAccessToken?icode=%s&isecret=%s", OPEN_URL, icode, isecret)

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var result RespData
	err = json.Unmarshal(data, &result)
	if err != nil {
		return "", err
	}
	if result.Code == 1 {
		return result.Data, nil
	} else {
		return "", errors.New(result.Msg)
	}
}

func uploadArt(token string, title string, keyword string, filename string, ispub int, islock int) error {
	if token == "" {
		return errors.New("token不能为空")
	}

	url := fmt.Sprintf("%s/upArt?token=%s", OPEN_URL, token)

	filecontent, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	snippet := string(filecontent)

	reqBody := new(bytes.Buffer)

	art := ArtReq{
		Title:   title,
		Keyword: keyword,
		Content: snippet,
		IsPub:   ispub,
		IsLock:  islock,
	}
	json.NewEncoder(reqBody).Encode(art)

	req, err := http.NewRequest("POST", url, reqBody)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var result RespMsg
	err = json.Unmarshal(data, &result)
	if err != nil {
		return err
	}

	if result.Code == 1 {
		return nil
	} else {
		return errors.New(result.Msg)
	}

}

// 构建 go build -ldflags '-s -w'
