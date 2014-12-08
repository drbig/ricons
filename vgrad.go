// See LICENSE.txt for licensing information.

// Vgrad generates simple vertical gradients.

package ricons

import (
	"math/rand"
	"time"
)

type vgradIconGen struct {
	r *rand.Rand
}

func (g *vgradIconGen) String() string {
	return "vgrad: simple vertical gradient"
}

func (g *vgradIconGen) NewIcon(width, height int) (*Icon, error) {
	i := NewIcon(width, height)

	re := g.r.Intn(255)
	sr := re / height
	if sr == 0 {
		sr = 1
	}
	if re > 128 {
		sr = -sr
	}
	if sr == 0 {
		sr = 1
	}
	gr := g.r.Intn(255)
	sg := gr / height
	if sg == 0 {
		sg = 1
	}
	if gr > 128 {
		sg = -sg
	}
	bl := g.r.Intn(255)
	sb := bl / height
	if bl > 128 {
		sb = -sb
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			idx := (y-i.Image.Rect.Min.Y)*i.Image.Stride + (x-i.Image.Rect.Min.X)*4
			i.Image.Pix[idx] = uint8(re)
			i.Image.Pix[idx+1] = uint8(gr)
			i.Image.Pix[idx+2] = uint8(bl)
			i.Image.Pix[idx+3] = 255
		}
		if re > 0 && re < 255 {
			re += sr
		}
		if gr > 0 && gr < 255 {
			gr += sg
		}
		if bl > 0 && bl < 255 {
			bl += sb
		}
	}
	return i, nil
}

func init() {
	g := &vgradIconGen{rand.New(rand.NewSource(time.Now().Unix()))}
	Register("vgrad", g)
}
