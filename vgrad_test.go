// See LICENSE.txt for licensing information.

package ricons

import "testing"

func TestVgradPresent(t *testing.T) {
	if gen, exist := Registry["vgrad"]; !exist {
		t.Fatal("couldn't find 'vgrad' generator")
	} else {
		g = gen
	}
}

func TestVgradStringer(t *testing.T) {
	if g.String() != "vgrad: simple vertical gradient" {
		t.Errorf("stringer mismatch: %v != %v", "vgrad: simple vertical gradient", g.String())
	}
}

func BenchmarkVgrad16x16(b *testing.B) {
	g, ok := Registry["vgrad"]
	if !ok {
		b.Fatal("couldn't find 'vgrad' generator")
	}
	ic := NewIcon(16, 16)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := g.Generate(ic); err != nil {
			b.Fatalf("(%d) error generating icon: %s", i+1, err)
		}
	}
}

func BenchmarkVgrad32x32(b *testing.B) {
	g, ok := Registry["vgrad"]
	if !ok {
		b.Fatal("couldn't find 'vgrad' generator")
	}
	ic := NewIcon(32, 32)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := g.Generate(ic); err != nil {
			b.Fatalf("(%d) error generating icon: %s", i+1, err)
		}
	}
}
