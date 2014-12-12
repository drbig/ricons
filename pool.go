// See LICENSE.txt for licensing information.

package ricons

import (
	"image"
	"sync"
)

type pool struct {
	size   int
	dim    image.Rectangle
	mtx    sync.Mutex
	images []*image.RGBA
}

// get retrieves an image from the pool or return a fresh one.
func (p *pool) get() *image.RGBA {
	var i *image.RGBA
	p.mtx.Lock()
	defer p.mtx.Unlock()
	l := len(p.images)
	if l > 0 {
		i, p.images = p.images[l-1], p.images[:l-1]
		return i
	}
	return image.NewRGBA(p.dim)
}

// put returns an image into the pool.
func (p *pool) put(i *image.RGBA) {
	p.mtx.Lock()
	defer p.mtx.Unlock()
	if len(p.images) >= p.size {
		return
	}
	p.images = append(p.images, i)
	return
}

// makePool initializes a new pool of images.
func makePool(width, height, size int) *pool {
	dim := image.Rect(0, 0, width, height)
	p := &pool{
		size:   size,
		dim:    dim,
		images: make([]*image.RGBA, size),
	}
	for i := 0; i < size; i++ {
		p.images[i] = image.NewRGBA(p.dim)
	}
	return p
}
