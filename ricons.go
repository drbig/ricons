// See LICENSE.txt for licensing information.

// Package ricons implements random icon generator framework.
package ricons

import (
	"errors"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
)

const (
	VERSION = `0.0.1` // framework version
)

// Format is an enum of available image formats.
type Format string

// EncoderOptions holds encoder options for supported output image formats.
type EncoderOptions struct {
	GIF  *gif.Options  // encoder options for GIF images
	JPEG *jpeg.Options // encoder options for JPEG images
}

// Icon holds an icon dimensions and the actual pixel data.
type Icon struct {
	Dim         image.Rectangle // dimensions
	Image       *image.RGBA     // actual pixel data
	EncoderOpts *EncoderOptions // encoder options
}

// Generator describes a minimal Icon generator.
type Generator interface {
	NewIcon(width, height int) (*Icon, error)
	fmt.Stringer
}

// Implemented image formats.
const (
	PNG  Format = "png"
	GIF  Format = "gif"
	JPEG Format = "jpeg"
)

// Error variables.
var (
	ErrUnknownFormat = errors.New("icon encode: unknown format")
)

// Global Generators registry.
var Registry = make(map[string]Generator, 0)

// Register registers a generator at compile time.
// It will panic if you try to overwrite an already existing Generator.
func Register(name string, g Generator) {
	if _, exists := Registry[name]; exists {
		panic(fmt.Sprintf("generator '%s' already registered", name))
	}
	Registry[name] = g
}

// Icon returns a basic fully initialised Icon.
// This includes default EncoderOptions.
func NewIcon(width, height int) *Icon {
	dim := image.Rect(0, 0, width, height)
	return &Icon{
		Dim:   dim,
		Image: image.NewRGBA(dim),
		EncoderOpts: &EncoderOptions{
			GIF:  &gif.Options{NumColors: 256},
			JPEG: &jpeg.Options{Quality: 75},
		},
	}
}

// Encode encodes and writes a given Icon in the given image format.
func (i *Icon) Encode(f Format, o io.Writer) error {
	switch f {
	case PNG:
		return png.Encode(o, i.Image)
	case GIF:
		return gif.Encode(o, i.Image, i.EncoderOpts.GIF)
	case JPEG:
		return jpeg.Encode(o, i.Image, i.EncoderOpts.JPEG)
	default:
		return ErrUnknownFormat
	}
}
