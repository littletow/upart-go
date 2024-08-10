package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
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
	OPEN_URL = "https://www.91demo.top/api/open"
)

// 获取token
func GetToken(icode string, isecret string) (string, error) {
	url := fmt.Sprintf("%s/vtoken?icode=%s&isecret=%s", OPEN_URL, icode, isecret)

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

// 上传文章
func UploadArt(token string, title string, keyword string, filename string, ispub int, islock int) error {
	if token == "" {
		return errors.New("token不能为空")
	}

	url := fmt.Sprintf("%s/upVArt?token=%s", OPEN_URL, token)

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
