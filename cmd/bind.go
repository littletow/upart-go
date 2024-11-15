package cmd

import (
	"bytes"
	"fmt"
	"image/color"
	"log"
	"math"
	"net"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "绑定账号，初始化配置文件",
	Long:  `绑定账号，初始化配置文件`,
	// Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// 连接Websocket，获取识别码
		// 调用ebtien代码，显示图片，让用户扫码，扫码后会进入mp，然后填写识别码，进行绑定。
		// 匹配成功后，写入到配置文件，并返回结果。
		// 绑定成功后，图片应用退出。结束绑定。
		fmt.Println("实现插入配置文件")
		ebiten.SetWindowSize(screenWidth, screenHeight)
		ebiten.SetWindowTitle("Text (Ebitengine Demo)")
		if err := ebiten.RunGame(&BindForm{}); err != nil {
			log.Fatal(err)
		}
	},
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

}

const (
	screenWidth  = 640
	screenHeight = 480
)

const sampleText = `  The quick brown fox jumps
over the lazy dog.`

var (
	mplusFaceSource *text.GoTextFaceSource
	mplusNormalFace *text.GoTextFace
	mplusBigFace    *text.GoTextFace
)

func init() {
	rootCmd.AddCommand(initCmd)
	s, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.MPlus1pRegular_ttf))
	if err != nil {
		log.Fatal(err)
	}
	mplusFaceSource = s

	mplusNormalFace = &text.GoTextFace{
		Source: mplusFaceSource,
		Size:   24,
	}
	mplusBigFace = &text.GoTextFace{
		Source: mplusFaceSource,
		Size:   32,
	}
}

type BindForm struct {
	glyphs      []text.Glyph
	showOrigins bool
}

func (g *BindForm) Update() error {
	// Initialize the glyphs for special (colorful) rendering.
	if len(g.glyphs) == 0 {
		op := &text.LayoutOptions{}
		op.LineSpacing = mplusNormalFace.Size * 1.5
		g.glyphs = text.AppendGlyphs(g.glyphs, sampleText, mplusNormalFace, op)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyO) {
		g.showOrigins = !g.showOrigins
	}
	return nil
}

func (g *BindForm) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Press O to show/hide origins")

	gray := color.RGBA{0x80, 0x80, 0x80, 0xff}

	{
		const x, y = 20, 20
		w, h := text.Measure(sampleText, mplusNormalFace, mplusNormalFace.Size*1.5)
		vector.DrawFilledRect(screen, x, y, float32(w), float32(h), gray, false)
		op := &text.DrawOptions{}
		op.GeoM.Translate(x, y)
		op.LineSpacing = mplusNormalFace.Size * 1.5
		text.Draw(screen, sampleText, mplusNormalFace, op)
	}
	{
		const x, y = 20, 120
		w, h := text.Measure(sampleText, mplusBigFace, mplusBigFace.Size*1.5)
		vector.DrawFilledRect(screen, x, y, float32(w), float32(h), gray, false)
		op := &text.DrawOptions{}
		op.GeoM.Translate(x, y)
		op.LineSpacing = mplusBigFace.Size * 1.5
		text.Draw(screen, sampleText, mplusBigFace, op)
	}
	{
		const x, y = 20, 220
		op := &text.DrawOptions{}
		op.GeoM.Rotate(math.Pi / 4)
		op.GeoM.Translate(x, y)
		op.Filter = ebiten.FilterLinear
		op.LineSpacing = mplusNormalFace.Size * 1.5
		text.Draw(screen, sampleText, mplusNormalFace, op)
	}
	{
		const x, y = 160, 220
		const lineSpacingInPixels = 80
		w, h := text.Measure(sampleText, mplusBigFace, lineSpacingInPixels)
		vector.DrawFilledRect(screen, x, y, float32(w), float32(h), gray, false)
		op := &text.DrawOptions{}
		// Add the width as the text rendering region's upper-right position comes to (0, 0)
		// when the horizontal alignment is right. The alignment is specified later (PrimaryAlign).
		op.GeoM.Translate(x+w, y)
		op.LineSpacing = lineSpacingInPixels
		// The primary alignment for the left-to-right direction is a horizontal alignment, and the end means the right.
		op.PrimaryAlign = text.AlignEnd
		text.Draw(screen, sampleText, mplusBigFace, op)
	}
	{
		const x, y = 240, 360
		op := &ebiten.DrawImageOptions{}
		// g.glyphs is initialized by text.AppendGlyphs.
		// You can customize how to render each glyph.
		// In this example, multiple colors are used to render glyphs.
		for i, gl := range g.glyphs {
			if gl.Image == nil {
				continue
			}
			op.GeoM.Reset()
			op.GeoM.Translate(x, y)
			op.GeoM.Translate(gl.X, gl.Y)
			op.ColorScale.Reset()
			r := float32(1)
			if i%3 == 0 {
				r = 0.5
			}
			g := float32(1)
			if i%3 == 1 {
				g = 0.5
			}
			b := float32(1)
			if i%3 == 2 {
				b = 0.5
			}
			op.ColorScale.Scale(r, g, b, 1)
			screen.DrawImage(gl.Image, op)
		}

		if g.showOrigins {
			for _, gl := range g.glyphs {
				vector.DrawFilledCircle(screen, x+float32(gl.OriginX), y+float32(gl.OriginY), 2, color.RGBA{0xff, 0, 0, 0xff}, true)
			}
		}
	}
}

func (g *BindForm) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
