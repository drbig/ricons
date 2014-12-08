// See LICENSE.txt for licensing information.

// Tests the basics of the framework on a concrete Generator.

package ricons

import (
	"bytes"
	"math/rand"
	"testing"
	"time"
)

var (
	g   Generator
	rng = rand.New(rand.NewSource(time.Now().Unix()))
)

func TestUniformPresent(t *testing.T) {
	if gen, exist := Registry["uniform"]; !exist {
		t.Fatal("couldn't find 'uniform' generator")
	} else {
		g = gen
	}
}

func TestUniformStringer(t *testing.T) {
	if g.String() != "uniform: single uniform color" {
		t.Errorf("stringer mismatch: %v != %v", "uniform: single uniform color", g.String())
	}
}

func TestUniformIcons(t *testing.T) {
	for i := 0; i < 5; i++ {
		icon, err := g.NewIcon(8+rng.Intn(24), 8+rng.Intn(24))
		if err != nil {
			t.Fatal(err)
		}
		c := icon.Image.At(rng.Intn(icon.Dim.Max.X), rng.Intn(icon.Dim.Max.Y))
		for j := 0; j < 5; j++ {
			ct := icon.Image.At(rng.Intn(icon.Dim.Max.X), rng.Intn(icon.Dim.Max.Y))
			if c != ct {
				t.Errorf("icon %d point %d color mismatch: %v != %v", i+1, j+1, ct, c)
			}
		}
	}
}

type encTest struct {
	f Format
	o []byte
}

func TestEncoders(t *testing.T) {
	var b bytes.Buffer
	cases := []encTest{
		encTest{PNG, []byte{137, 80, 78, 71, 13, 10, 26, 10}},
		encTest{GIF, []byte{71, 73, 70, 56}},
		encTest{JPEG, []byte{0xff, 0xd8, 0xff}},
	}
	icon, err := g.NewIcon(8+rng.Intn(24), 8+rng.Intn(24))
	if err != nil {
		t.Fatal(err)
	}

	for i, c := range cases {
		b.Reset()
		if err := icon.Encode(c.f, &b); err != nil {
			t.Errorf("(%d) icon encode to %v failed: %s", i+1, c.f, err)
		} else {
			if !bytes.Equal(b.Bytes()[:len(c.o)], c.o) {
				t.Errorf("(%d) icon encode bad header for %v", i+1, c.f)
			}
		}
	}
}

func TestBadEncoder(t *testing.T) {
	var b bytes.Buffer
	icon, err := g.NewIcon(8+rng.Intn(24), 8+rng.Intn(24))
	if err != nil {
		t.Fatal(err)
	}

	if err := icon.Encode(Format("DefinitelyUnknownFormat"), &b); err == nil {
		t.Errorf("didn't err on not defined image encoding")
	}
}

func TestBadRegister(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("didn't panic on already registered generator")
		}
	}()
	Register("uniform", g)
}

func BenchmarkUniform16x16(b *testing.B) {
	g, ok := Registry["uniform"]
	if !ok {
		b.Fatal("couldn't find 'uniform' generator")
	}
	for i := 0; i < b.N; i++ {
		_, err := g.NewIcon(16, 16)
		if err != nil {
			b.Fatalf("(%d) error generating icon: %s", i+1, err)
		}
	}
}

func BenchmarkUniform32x32(b *testing.B) {
	g, ok := Registry["uniform"]
	if !ok {
		b.Fatal("couldn't find 'uniform' generator")
	}
	for i := 0; i < b.N; i++ {
		_, err := g.NewIcon(32, 32)
		if err != nil {
			b.Fatalf("(%d) error generating icon: %s", i+1, err)
		}
	}
}
