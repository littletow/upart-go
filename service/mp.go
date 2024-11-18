package service

import (
	"bytes"
	"gart/fonts"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	mpPageWidth  = 320
	mpPageHeight = 320
)

var (
	mpImage *ebiten.Image
)

func init() {
	img, _, err := image.Decode(bytes.NewReader(fonts.VisitCode))
	if err != nil {
		log.Fatal(err)
	}
	mpImage = ebiten.NewImageFromImage(img)
}

// 显示Mp小程序码
func ShowMpCode() {
	mp := MPPage{}
	ebiten.SetWindowSize(mpPageWidth, mpPageHeight)
	ebiten.SetWindowTitle("豆子碎片小程序码")
	if err := ebiten.RunGame(&mp); err != nil {
		log.Fatal(err)
	}
}

type MPPage struct {
}

func (g *MPPage) Update() error {
	return nil
}

func (g *MPPage) Draw(screen *ebiten.Image) {
	screen.Fill(backgroundColor)
	{
		const x, y = 20, 20
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(0.65, 0.65)
		op.GeoM.Translate(x, y)
		screen.DrawImage(mpImage, op)
	}

}

func (g *MPPage) Layout(outsideWidth, outsideHeight int) (int, int) {
	return mpPageWidth, mpPageHeight
}
