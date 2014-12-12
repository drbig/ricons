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
	ic := NewIcon(16, 16)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := g.Generate(ic); err != nil {
			b.Fatalf("(%d) error generating icon: %s", i+1, err)
		}
	}
}

func BenchmarkGrid32x32(b *testing.B) {
	g, ok := Registry["grid"]
	if !ok {
		b.Fatal("couldn't find 'grid' generator")
	}
	ic := NewIcon(32, 32)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := g.Generate(ic); err != nil {
			b.Fatalf("(%d) error generating icon: %s", i+1, err)
		}
	}
}
