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
	Title     string `json:"title"`
	Desc      string `json:"desc"`
	Tags      int    `json:"tags"`
	Github    string `json:"github"`
	Snippet   string `json:"snippet"`
	LockState int    `json:"lockstate"`
	AType     int    `json:"atype"`
}

const (
	OPEN_URL = "https://open.91demo.top/open"
)

var (
	artType   int // 1 snippet 2 lib 3 art
	title     string
	desc      string
	filename  string
	github    string
	tag       int // 1 rust 2 go 3 mp 4 web 5 sql 6 dev
	lockstate int // 1 lock

)

func init() {
	flag.IntVar(&artType, "a", 1, "upload article type")
	flag.StringVar(&title, "b", "", "article title")
	flag.StringVar(&desc, "d", "", "article desc")
	flag.StringVar(&filename, "f", "", "content filename")
	flag.StringVar(&github, "g", "", "github url")
	flag.IntVar(&tag, "t", 6, "article tag")
	flag.IntVar(&lockstate, "l", 1, "article lockstate")
}

func printHelp() {
	fmt.Println("./upart.exe -a 1 -b title -d desc -f filename -g github -t tag -l lockstate")
}

func main() {
	flag.Parse()
	var err error

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println("Fail to get os dir,", err)
		os.Exit(1)
	}
	confFile := filepath.Join(dir, "conf.ini")
	cfg, err := ini.Load(confFile)
	if err != nil {
		fmt.Println("Fail to read conf file,", err)
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
			fmt.Println("Fail to get token,", err)
			os.Exit(1)
		}
		expireAtStr := strconv.FormatInt(now+7000, 10)
		cfg.Section("User").Key("token").SetValue(token)
		cfg.Section("User").Key("expire_at").SetValue(expireAtStr)
		cfg.SaveTo("conf.ini")
	}
	if token == "" || artType == 0 || title == "" || desc == "" {
		printHelp()
		fmt.Println("Please provide upload article parameter")
		os.Exit(1)
	}
	// 上传文章
	err = uploadArt(token, artType, title, desc, filename, tag, lockstate)
	if err != nil {
		fmt.Println("Fail to upload article,", err)
		os.Exit(1)
	}

	fmt.Println("Success to upload article")
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

func uploadArt(token string, artType int, title string, desc string, filename string, tag int, lockstate int) error {
	if token == "" {
		return errors.New("token is empty")
	}

	url := fmt.Sprintf("%s/upArt?token=%s", OPEN_URL, token)

	if artType == 1 {
		if filename == "" {
			return errors.New("filename is empty")
		}
	} else if artType == 2 {
		if github == "" {
			return errors.New("github is empty")
		}
	} else if artType == 3 {
		if filename == "" {
			return errors.New("filename is empty")
		}
		url = fmt.Sprintf("%s/upArtL1?token=%s", OPEN_URL, token)
	}

	snippet := ""

	if filename != "" {
		filecontent, err := os.ReadFile(filename)
		if err != nil {
			return err
		}
		snippet = string(filecontent)
	}

	art := ArtReq{
		Title:     title,
		Desc:      desc,
		Tags:      tag,
		Github:    github,
		Snippet:   snippet,
		AType:     artType,
		LockState: lockstate,
	}

	reqBody := new(bytes.Buffer)
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
