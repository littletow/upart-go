package service

import (
	"bytes"
	"gart/fonts"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	adPageWidth  = 320
	adPageHeight = 320
)

var (
	adImage *ebiten.Image
)

func init() {
	img, _, err := image.Decode(bytes.NewReader(fonts.VisitAdCode))
	if err != nil {
		log.Fatal(err)
	}
	adImage = ebiten.NewImageFromImage(img)
}

// 显示AD小程序码
func ShowAdCode() {
	ad := AdPage{}
	ebiten.SetWindowSize(adPageWidth, adPageHeight)
	ebiten.SetWindowTitle("豆子碎片广告页面")
	if err := ebiten.RunGame(&ad); err != nil {
		log.Fatal(err)
	}
}

type AdPage struct {
}

func (g *AdPage) Update() error {
	return nil
}

func (g *AdPage) Draw(screen *ebiten.Image) {
	screen.Fill(backgroundColor)
	{
		const x, y = 20, 20
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(0.65, 0.65)
		op.GeoM.Translate(x, y)
		screen.DrawImage(adImage, op)
	}

}

func (g *AdPage) Layout(outsideWidth, outsideHeight int) (int, int) {
	return adPageWidth, adPageHeight
}
