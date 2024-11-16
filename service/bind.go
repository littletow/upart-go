package service

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"gart/fonts"
	"gart/utils"
	"image"
	"image/color"
	"log"
	"math"
	"time"

	"github.com/gorilla/websocket"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Maximum message size allowed from peer.
	maxMessageSize = 512

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10
)

const (
	screenWidth  = 320
	screenHeight = 480
)

const (
	logoText  = `豆子碎片`
	macText   = `Mac地址：`
	vcodeText = `验证码：`
	helpText  = `请使用微信扫描下方小程序码，
然后输入验证码即可绑定账号。`
)

var (
	content1 string
	content2 string
)

var (
	visit           *ebiten.Image
	faceSource      *text.GoTextFaceSource
	smallFace       *text.GoTextFace
	normalFace      *text.GoTextFace
	bigFace         *text.GoTextFace
	backgroundColor = color.RGBA{0xfa, 0xf8, 0xef, 0xff}
	gray            = color.RGBA{0x80, 0x80, 0x80, 0xff}
	frameColor      = color.RGBA{0xbb, 0xad, 0xa0, 0xff}
)

func init() {
	img, _, err := image.Decode(bytes.NewReader(fonts.VisitCode))
	if err != nil {
		log.Fatal(err)
	}
	visit = ebiten.NewImageFromImage(img)
	s, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.AlibabaLight_ttf))
	if err != nil {
		log.Fatal(err)
	}
	faceSource = s
	smallFace = &text.GoTextFace{
		Source: faceSource,
		Size:   18,
	}
	normalFace = &text.GoTextFace{
		Source: faceSource,
		Size:   20,
	}
	bigFace = &text.GoTextFace{
		Source: faceSource,
		Size:   22,
	}
}

// 命令消息
type CmdMsg struct {
	UUID    string `json:"uuid"`    // 唯一码，识别发送给谁？
	Cmd     int    `json:"cmd"`     // 命令
	Content string `json:"content"` // 内容，约定好顺序后，以逗号分割
}

// 读取ws
func WsRead(done chan struct{}, conn *websocket.Conn, form *BindForm) {
	defer close(done)
	conn.SetReadLimit(maxMessageSize)
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			return
		}
		log.Printf("recv: %s", message)
		var cmdMsg CmdMsg
		err = json.Unmarshal(message, &cmdMsg)
		if err != nil {
			log.Println("消息解析错误，", err)
			continue
		}
		if cmdMsg.UUID != form.Mac {
			log.Println("不是发给自己的消息")
			continue
		}
		switch cmdMsg.Cmd {
		case 1000:
			form.Vcode = cmdMsg.Content
			log.Println("验证码，", cmdMsg.Content)
		case 1001:
			content := cmdMsg.Content
			log.Println("识别码，", content)
		}
	}
}

func WsWrite(ctx context.Context, done chan struct{}, conn *websocket.Conn) {
	defer conn.Close()
	ticker := time.NewTicker(pingPeriod)
	defer ticker.Stop()
	for {
		select {
		case <-done:
			return
		case <-ticker.C:
			if err := conn.WriteControl(websocket.PingMessage, []byte{}, time.Now().Add(writeWait)); err != nil {
				log.Println("ping:", err)
				return
			}
		case <-ctx.Done():
			log.Println("interrupt")
			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}

// 连接Websocket
func WsConn(ctx context.Context, form *BindForm) error {
	// 获取识别码，并接收绑定结果
	// 连接Websocket
	mac := utils.GetMacAddr()
	if mac == "" {
		return errors.New("无法获取MAC地址")
	}
	vcode := utils.GenVcode(mac)
	url := fmt.Sprintf("ws://www.91demo.top:37021/bean?mac=%s&vcode=%s", mac, vcode)

	c, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return err
	}

	form.Mac = mac
	done := make(chan struct{})
	go WsWrite(ctx, done, c)
	go WsRead(done, c, form)

	return nil
}

// 显示图片
func ShowImage() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// 使用ebtien显示
	// isEnable := viper.GetBool("is_enable")
	// fmt.Println("isEnable,", isEnable)
	form := BindForm{}
	err := WsConn(ctx, &form)
	if err != nil {
		log.Fatal(err)
	}

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("绑定豆子碎片账号")
	if err := ebiten.RunGame(&form); err != nil {
		log.Fatal(err)
	}
}

type BindForm struct {
	Mac   string
	Vcode string
}

func (g *BindForm) Update() error {
	return nil
}

func (g *BindForm) Draw(screen *ebiten.Image) {
	// ebitenutil.DebugPrint(screen, "Press O to show/hide origins")
	// 上面写上MAC地址和验证码，下面写使用方法，然后右下角写Logo，斜角45度。
	screen.Fill(backgroundColor)
	vector.DrawFilledRect(screen, 5, 5, float32(310), float32(150), gray, false)
	vector.DrawFilledRect(screen, 5, 165, float32(310), float32(310), frameColor, false)
	{
		const x, y = 20, 20
		content1 = fmt.Sprintf("%s%s", macText, g.Mac)
		// fmt.Println("content1,", content1)
		w, h := text.Measure(content1, normalFace, normalFace.Size*1)
		vector.DrawFilledRect(screen, x, y, float32(w), float32(h), gray, false)
		op := &text.DrawOptions{}
		op.GeoM.Translate(x, y)
		op.LineSpacing = normalFace.Size * 1
		text.Draw(screen, content1, normalFace, op)
	}
	{
		const x, y = 20, 60
		w, h := text.Measure(helpText, smallFace, smallFace.Size*1.5)
		vector.DrawFilledRect(screen, x, y, float32(w), float32(h), gray, false)
		op := &text.DrawOptions{}
		op.GeoM.Translate(x, y)
		op.LineSpacing = smallFace.Size * 1.5
		text.Draw(screen, helpText, smallFace, op)
	}
	{
		const x, y = 20, 120
		content2 = fmt.Sprintf("%s%s", vcodeText, g.Vcode)
		// fmt.Println("content2,", content2)
		w, h := text.Measure(content2, normalFace, normalFace.Size*1)
		vector.DrawFilledRect(screen, x, y, float32(w), float32(h), gray, false)
		op := &text.DrawOptions{}
		op.GeoM.Translate(x, y)
		op.LineSpacing = normalFace.Size * 1
		text.Draw(screen, content2, normalFace, op)
	}
	{
		const x, y = 240, 140

		op := &text.DrawOptions{}
		op.GeoM.Rotate(2*math.Pi - math.Pi/4)
		op.GeoM.Translate(x, y)
		op.Filter = ebiten.FilterLinear
		op.LineSpacing = bigFace.Size * 1.5
		text.Draw(screen, logoText, bigFace, op)
	}
	{
		const x, y = 20, 180
		// pw, ph := visit.Bounds().Dx(), visit.Bounds().Dy()
		// fmt.Println("pw,ph", pw, ph)
		op := &ebiten.DrawImageOptions{}
		// op.GeoM.Translate(-float64(pw)/2, -float64(ph)/2)
		op.GeoM.Scale(0.65, 0.65)
		op.GeoM.Translate(x, y)

		screen.DrawImage(visit, op)
	}

}

func (g *BindForm) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
