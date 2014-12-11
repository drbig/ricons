// See LICENSE.txt for licensing information.

// Square-based symmetric icons.

package ricons

import (
	"image"
	"image/color"
	"image/draw"
)

type symsquareIconGen struct {
	s int
	p []color.Color
	b image.Image
}

func (g *symsquareIconGen) String() string {
	return "symsquare: symmetric square-based icons"
}

func (g *symsquareIconGen) NewIcon(width, height int) (*Icon, error) {
	i := NewIcon(width, height)
	draw.Draw(i.Image, image.Rect(0, 0, width, height), g.b, image.ZP, draw.Src)

	cx := width / 2
	cy := height / 2
	c := image.NewUniform(g.p[<-chRnd%len(g.p)])

	for cnt := 0; cnt < 8; cnt++ {
		s := <-chRnd % (cx / 2)
		x := <-chRnd % (cx - s)
		y := <-chRnd % (cy - s)

		r := image.Rect(x, y, x+s, y+s)
		draw.Draw(i.Image, r, c, image.ZP, draw.Src)

		r = image.Rect(width-x-s, y, (width - x), y+s)
		draw.Draw(i.Image, r, c, image.ZP, draw.Src)

		r = image.Rect(x, height-y-s, x+s, (height - y))
		draw.Draw(i.Image, r, c, image.ZP, draw.Src)

		r = image.Rect(width-x-s, height-y-s, (width - x), (height - y))
		draw.Draw(i.Image, r, c, image.ZP, draw.Src)
	}

	return i, nil
}

func init() {
	g := &symsquareIconGen{
		5,
		[]color.Color{
			color.RGBA{0xaa, 0x66, 0x66, 0xff},
			color.RGBA{0x66, 0xaa, 0x66, 0xff},
			color.RGBA{0x66, 0x66, 0xaa, 0xff},
		},
		image.NewUniform(color.RGBA{0xff, 0xff, 0xff, 0xff}),
	}
	Register("symsquare", g)
}
