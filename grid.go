// See LICENSE.txt for licensing information.

// Grid generates grid-based icons.

package ricons

import (
	"image"
	"image/color"
	"image/draw"
	"math/rand"
	"time"
)

type gridIconGen struct {
	r *rand.Rand
	s int
	p []color.Color
	b image.Image
}

func (g *gridIconGen) String() string {
	return "grid: grid-based icons"
}

func (g *gridIconGen) NewIcon(width, height int) (*Icon, error) {
	i := NewIcon(width, height)
	r := image.Rect(0, 0, width, height)
	draw.Draw(i.Image, r, g.b, image.ZP, draw.Src)

	gr := make([][]bool, g.s)
	el := make([]bool, g.s*g.s)
	for i := range gr {
		gr[i] = el[i*g.s : (i+1)*g.s]
	}

	for i := 0; i < 8; i++ {
		x := g.r.Intn(g.s)
		y := g.r.Intn(g.s)
		gr[x][y] = true
		if g.r.Float32() > 0.5 {
			gr[g.s-x-1][y] = true
		}
		if g.r.Float32() > 0.5 {
			gr[x][g.s-y-1] = true
		}
	}

	c := image.NewUniform(g.p[g.r.Intn(len(g.p))])
	w := width / g.s
	h := height / g.s
	for y := 0; y < g.s; y++ {
		for x := 0; x < g.s; x++ {
			if gr[x][y] {
				r.Min.X = x * w
				r.Min.Y = y * h
				r.Max.X = (x + 1) * w
				r.Max.Y = (y + 1) * h
				draw.Draw(i.Image, r, c, image.ZP, draw.Src)
			}
		}
	}

	return i, nil
}

func init() {
	g := &gridIconGen{
		rand.New(rand.NewSource(time.Now().Unix())),
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
