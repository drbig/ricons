// See LICENSE.txt for licensing information.

package ricons

import "testing"

func TestGridPresent(t *testing.T) {
	if gen, exist := Registry["grid"]; !exist {
		t.Fatal("couldn't find 'grid' generator")
	} else {
		g = gen
	}
}

func TestGridStringer(t *testing.T) {
	if g.String() != "grid: grid-based icons" {
		t.Errorf("stringer mismatch: %v != %v", "grid: grid-based icons", g.String())
	}
}

func BenchmarkGrid16x16(b *testing.B) {
	g, ok := Registry["grid"]
	if !ok {
		b.Fatal("couldn't find 'grid' generator")
	}
	for i := 0; i < b.N; i++ {
		_, err := g.NewIcon(16, 16)
		if err != nil {
			b.Fatalf("(%d) error generating icon: %s", i+1, err)
		}
	}
}

func BenchmarkGrid32x32(b *testing.B) {
	g, ok := Registry["grid"]
	if !ok {
		b.Fatal("couldn't find 'grid' generator")
	}
	for i := 0; i < b.N; i++ {
		_, err := g.NewIcon(32, 32)
		if err != nil {
			b.Fatalf("(%d) error generating icon: %s", i+1, err)
		}
	}
}
