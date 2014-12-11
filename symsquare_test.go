// See LICENSE.txt for licensing information.

// Tests the basics of the framework on a concrete Generator.

package ricons

import "testing"

func TestSymsquarePresent(t *testing.T) {
	if gen, exist := Registry["symsquare"]; !exist {
		t.Fatal("couldn't find 'symsquare' generator")
	} else {
		g = gen
	}
}

func TestSymsquareStringer(t *testing.T) {
	if g.String() != "symsquare: symmetric square-based icons" {
		t.Errorf("stringer mismatch: %v != %v", "symsquare: symmetric square-based icons", g.String())
		return
	}
}

func BenchmarkSymsquare16x16(b *testing.B) {
	g, ok := Registry["symsquare"]
	if !ok {
		b.Fatal("couldn't find 'symsquare' generator")
	}
	ic := NewIcon(16, 16)
	iconPool[ic.Image.Bounds()] = ic
	for i := 0; i < b.N; i++ {
		_, err := g.NewIcon(16, 16)
		if err != nil {
			b.Fatalf("(%d) error generating icon: %s", i+1, err)
		}
	}
}

func BenchmarkSymsquare32x32(b *testing.B) {
	g, ok := Registry["symsquare"]
	if !ok {
		b.Fatal("couldn't find 'symsquare' generator")
	}
	ic := NewIcon(32, 32)
	iconPool[ic.Image.Bounds()] = ic
	for i := 0; i < b.N; i++ {
		_, err := g.NewIcon(32, 32)
		if err != nil {
			b.Fatalf("(%d) error generating icon: %s", i+1, err)
		}
	}
}
