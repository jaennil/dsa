package bimap

import (
	"testing"
)

func TestBiMap(t *testing.T) {
	bm := NewBiMap[rune, string]()
	bm.Insert('r', "rune")
	if v, exists := bm.Get('r'); exists {
		println(v)
	} else {
		println("not exists")
	}
	if v, exists := bm.GetInverse("rune"); exists {
		println(string(v))
	} else {
		println("not exists")
	}
}
