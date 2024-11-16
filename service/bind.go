package service

import (
	"bytes"
	"fmt"
	"gart/fonts"
	"image/color"
	"log"
	"net"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/spf13/viper"
)

const (
	screenWidth  = 320
	screenHeight = 480
)

const sampleText1 = `Mac地址：`
const sampleText2 = `验证码：`
const sampleText3 = `04D4C404CBE2`
const sampleText4 = `aZx86782`

var (
	faceSource      *text.GoTextFaceSource
	normalFace      *text.GoTextFace
	bigFace         *text.GoTextFace
	backgroundColor = color.RGBA{0xfa, 0xf8, 0xef, 0xff}
	gray            = color.RGBA{0x80, 0x80, 0x80, 0xff}
	frameColor      = color.RGBA{0xbb, 0xad, 0xa0, 0xff}
)

func init() {

	s, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.AlibabaLight_ttf))
	if err != nil {
		log.Fatal(err)
	}
	faceSource = s

	normalFace = &text.GoTextFace{
		Source: faceSource,
		Size:   20,
	}
	bigFace = &text.GoTextFace{
		Source: faceSource,
		Size:   24,
	}
}

// 获取本机MAC地址
func GetMacAddr() string {
	ints, err := net.Interfaces()
	if err != nil {
		panic("无法获取本机网络，" + err.Error())
	}

	i0 := ints[0]
	mac := i0.HardwareAddr.String()
	return mac
}

// 连接Websocket
func ConnWs() {
	// 获取识别码，并接收绑定结果

}

// 显示图片
func ShowImage() {
	// 使用ebtien显示
	isEnable := viper.GetBool("is_enable")
	fmt.Println("isEnable,", isEnable)
	mac := GetMacAddr()
	fmt.Println("mac,", mac)
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Text (Ebitengine Demo)")
	if err := ebiten.RunGame(&BindForm{}); err != nil {
		log.Fatal(err)
	}
}

type BindForm struct {
	Vcode string
	Mac   string
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
		w1, h1 := text.Measure(sampleText1, normalFace, normalFace.Size*1)
		w2, h2 := text.Measure(sampleText3, bigFace, bigFace.Size*1)
		w := w1 + w2 + 10
		h := h2
		y1 := y + (h2 - h1)
		y2 := y + (h2-h1)/2
		vector.DrawFilledRect(screen, x, y, float32(w), float32(h), gray, false)
		op := &text.DrawOptions{}
		op.GeoM.Translate(x+5, y1)
		op.LineSpacing = normalFace.Size * 1
		text.Draw(screen, sampleText1, normalFace, op)
		op = &text.DrawOptions{}
		op.GeoM.Translate(x+5+w1, y2)
		op.LineSpacing = bigFace.Size * 1
		text.Draw(screen, sampleText3, bigFace, op)
	}
	{
		const x, y = 20, 120
		w, h := text.Measure(sampleText2, normalFace, normalFace.Size*1)
		vector.DrawFilledRect(screen, x, y, float32(w), float32(h), gray, false)
		op := &text.DrawOptions{}
		op.GeoM.Translate(x, y)
		op.LineSpacing = normalFace.Size * 1
		text.Draw(screen, sampleText2, normalFace, op)
	}

}

func (g *BindForm) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
