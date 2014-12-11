// See LICENSE.txt for licensing information.

// Grid generates grid-based icons.

package ricons

import (
	"image"
	"image/color"
	"image/draw"
)

type gridIconGen struct {
	s int
	p []color.Color
	b image.Image
}

func (g *gridIconGen) String() string {
	return "grid: grid-based icons"
}

func (g *gridIconGen) NewIcon(width, height int) (*Icon, error) {
	i := NewIcon(width, height)
	draw.Draw(i.Image, image.Rect(0, 0, width, height), g.b, image.ZP, draw.Src)

	gr := make([][]bool, g.s)
	el := make([]bool, g.s*g.s)
	for i := range gr {
		gr[i] = el[i*g.s : (i+1)*g.s]
	}

	for i := 0; i < 8; i++ {
		x := <-chRnd % g.s
		y := <-chRnd % g.s
		gr[x][y] = true
		if <-chRnd > 126 {
			gr[g.s-x-1][y] = true
		}
		if <-chRnd > 126 {
			gr[x][g.s-y-1] = true
		}
	}

	c := image.NewUniform(g.p[<-chRnd%len(g.p)])
	w := width / g.s
	h := height / g.s
	for y := 0; y < g.s; y++ {
		for x := 0; x < g.s; x++ {
			if gr[x][y] {
				r := image.Rect(x*w, y*h, (x+1)*w, (y+1)*h)
				draw.Draw(i.Image, r, c, image.ZP, draw.Src)
			}
		}
	}

	return i, nil
}

func init() {
	g := &gridIconGen{
		5,
		[]color.Color{
			color.RGBA{0xaa, 0x66, 0x66, 0xff},
			color.RGBA{0x66, 0xaa, 0x66, 0xff},
			color.RGBA{0x66, 0x66, 0xaa, 0xff},
		},
		image.NewUniform(color.RGBA{0xdd, 0xdd, 0xdd, 0xff}),
	}
	Register("grid", g)
}
