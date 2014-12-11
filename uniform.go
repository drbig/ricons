// See LICENSE.txt for licensing information.

// Uniform generates an uniformly colored icon.
// Useful for testing, not very interesting in practice.

package ricons

import (
	"image"
	"image/color"
	"image/draw"
)

type uniformIconGen struct {
}

func (g *uniformIconGen) String() string {
	return "uniform: single uniform color"
}

func (g *uniformIconGen) NewIcon(width, height int) (*Icon, error) {
	i := NewIcon(width, height)
	re := uint8(<-chRnd)
	gr := uint8(<-chRnd)
	bl := uint8(<-chRnd)
	bg := image.NewUniform(&color.RGBA{re, gr, bl, 0xff})
	draw.Draw(i.Image, i.Dim, bg, image.ZP, draw.Src)
	return i, nil
}

func init() {
	g := &uniformIconGen{}
	Register("uniform", g)
}
