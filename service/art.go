package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
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

type UploadArtReq struct {
	Title   string `json:"title"`
	Keyword string `json:"keyword"`
	Content string `json:"content"`
	IsLock  int    `json:"islock"`
	IsPub   int    `json:"ispub"`
}

type RemoveArtReq struct {
	Uuid string `json:"uuid"`
}

type ArtItem struct {
	Id         int64     `json:"id"`         // ID
	Openid     string    `json:"openid"`     // 微信openid
	Uuid       string    `json:"uuid"`       // 是标识资料的唯一值，与内容进行关联，不用id是防止黑客迭代获取资料。
	Title      string    `json:"title"`      // 标题，显示在首页，有字数限制。
	Keyword    string    `json:"keyword"`    // 关键字，用于搜索
	IsPub      int       `json:"ispub"`      // 是否公开，默认不公开 1 公开，需要审核
	IsLock     int       `json:"islock"`     // 是否加锁 1 加锁，需要公开的情况才能加锁。
	Views      int       `json:"views"`      // 是浏览次数，无论是否加锁，或者公开。首页排序时使用。
	Status     int       `json:"status"`     // 状态 1 正常展示 2 审核中 3 审核被拒 4 封禁
	Createtime time.Time `json:"createtime"` // 创建时间
	Updatetime time.Time `json:"updatetime"` // 更新时间
}

type SearchArtRsp struct {
	Code int       `json:"code"`
	Msg  string    `json:"msg"`
	Data []ArtItem `json:"data"`
}

type AreaRsp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data string `json:"data"`
}

const (
	SHARE_KEY = "20!I@LOVE$CHINA#24"
	OPEN_URL  = "https://www.91demo.top/api/open"
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

// 获取文章内容
func GetFileContent(filename string) ([]byte, error) {
	filecontent, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return filecontent, nil
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

	art := UploadArtReq{
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

// 根据用户输入内容查找文章，通过标题和关键字模糊查询
func SearchArt(token string, content string) ([]ArtItem, error) {
	if token == "" {
		return nil, errors.New("token不能为空")
	}

	url := fmt.Sprintf("%s/qVArt?token=%s&kw=%s", OPEN_URL, token, content)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result SearchArtRsp
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}

	if result.Code == 1 {
		return result.Data, nil
	} else {
		return nil, errors.New(result.Msg)
	}
}

// 删除文章
func RemoveArt(token string, uuid string) error {
	if token == "" {
		return errors.New("token不能为空")
	}

	url := fmt.Sprintf("%s/rmVArt?token=%s", OPEN_URL, token)

	reqBody := new(bytes.Buffer)
	art := RemoveArtReq{
		Uuid: uuid,
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

// 更新文章请求
type UpdateArtReq struct {
	Uuid     string `json:"uuid"`     // uuid
	Title    string `json:"title"`    // 题目
	Keyword  string `json:"keyword"`  // 关键字
	Content  string `json:"content"`  // 内容
	IsPub    int    `json:"ispub"`    // 是否公开
	IsLock   int    `json:"islock"`   // 是否加锁
	AreaFlag int    `json:"areaflag"` // 区域限制标志，0 不限制 1 限制国家 2 限制省份 3 限制城市
	AreaCont string `json:"areacont"` // 区域内容
	UptType  int    `json:"utype"`    // 更新类型 1，题目 2，关键字 3，内容 4，是否公开 5，是否加锁 6，同城 7，强制公开文章
}

// 更新文章信息
func UpdateArt(token string, uar *UpdateArtReq) error {
	if token == "" {
		return errors.New("token不能为空")
	}

	url := fmt.Sprintf("%s/uptVArt?token=%s", OPEN_URL, token)

	reqBody := new(bytes.Buffer)

	json.NewEncoder(reqBody).Encode(uar)

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

// 获取有效省份
func GetValidProvince(token string) (string, error) {
	if token == "" {
		return "", errors.New("token不能为空")
	}

	url := fmt.Sprintf("%s/getVValidProvince?token=%s", OPEN_URL, token)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var result AreaRsp
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

// 获取有效城市
func GetValidCity(token string, province string) (string, error) {
	if token == "" {
		return "", errors.New("token不能为空")
	}

	url := fmt.Sprintf("%s/getVValidCity?token=%s&province=%s", OPEN_URL, token, province)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var result AreaRsp
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
