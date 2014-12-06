// Uniform generates an uniformly colored icon.
// Useful for testing, not very interesting in practice.

package ricons

import (
	"image"
	"image/color"
	"image/draw"
	"math/rand"
	"time"
)

type uniformIconGen struct {
	r *rand.Rand
}

func (g *uniformIconGen) String() string {
	return "uniform: icon of uniform color"
}

func (g *uniformIconGen) NewIcon(width, height int) (*Icon, error) {
	i := NewIcon(width, height)
	v := uint8(g.r.Intn(255))
	bg := image.NewUniform(&color.RGBA{v, v, v, 0xff})
	draw.Draw(i.Image, i.Dim, bg, image.ZP, draw.Src)
	return i, nil
}

func init() {
	g := &uniformIconGen{rand.New(rand.NewSource(time.Now().Unix()))}
	Register("uniform", g)
}
