// See LICENSE.txt for licensing information.

// Square-based symmetric icons.

package ricons

import (
	"image"
	"image/color"
	"image/draw"
	"math/rand"
	"time"
)

type symsquareIconGen struct {
	r *rand.Rand
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
	c := image.NewUniform(g.p[g.r.Intn(len(g.p))])

	for cnt := 0; cnt < 8; cnt++ {
		s := g.r.Intn(cx / 2)
		x := g.r.Intn(cx - s)
		y := g.r.Intn(cy - s)

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
		rand.New(rand.NewSource(time.Now().Unix())),
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
