package ricons

import (
  "testing"
  "math/rand"
  "time"
)

var (
  g Generator
  rng = rand.New(rand.NewSource(time.Now().Unix()))
)

func TestHasUniformGen(t *testing.T) {
  if gen, exist := Registry["uniform"]; !exist {
    t.Fatal("Couldn't find 'uniform' generator")
  } else {
    g = gen
  }
}

func TestFiveRandomIcons(t *testing.T) {
  for i := 0; i < 5; i++ {
    icon, err := g.NewIcon(8 + rng.Intn(24), 8 + rng.Intn(23))
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
