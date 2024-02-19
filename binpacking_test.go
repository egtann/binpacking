package binpacking

import (
	"fmt"
	"reflect"

	"testing"
)

func BenchmarkPack(b *testing.B) {
	allBoxes := []Box[string]{
		{Width: 220, Height: 160, Depth: 100, Weight: 110, Item: "Box1"},
		{Width: 260, Height: 145, Depth: 145, Weight: 120, Item: "Box2"},
		{Width: 270, Height: 185, Depth: 110, Weight: 140, Item: "Box3"},
		{Width: 310, Height: 220, Depth: 140, Weight: 210, Item: "Box4"},
		{Width: 300, Height: 210, Depth: 200, Weight: 250, Item: "Box5"},
		{Width: 300, Height: 300, Depth: 130, Weight: 290, Item: "Box6"},
		{Width: 370, Height: 270, Depth: 150, Weight: 300, Item: "Box7"},
		{Width: 300, Height: 300, Depth: 250, Weight: 360, Item: "Box8"},
		{Width: 470, Height: 280, Depth: 210, Weight: 400, Item: "Box9"},
		{Width: 430, Height: 315, Depth: 200, Weight: 430, Item: "Box10"},
		{Width: 330, Height: 330, Depth: 350, Weight: 500, Item: "Box11"},
		{Width: 465, Height: 350, Depth: 370, Weight: 650, Item: "Box12"},
	}
	items := []Item{
		goods{20, 100, 30},
		goods{100, 20, 30},
		goods{20, 100, 30},
		goods{100, 20, 30},
		goods{100, 20, 30},
		goods{100, 100, 30},
		goods{100, 100, 30},
		goods{100, 100, 30},
		goods{100, 100, 30},
		goods{100, 100, 30},
		goods{100, 100, 30},
		goods{100, 100, 30},
		goods{100, 100, 30},
		goods{100, 100, 30},
	}
	for i := 0; i < b.N; i++ {
		_, err := Pack(allBoxes, items)
		if err != nil {
			b.Error(err)
		}
	}
}

type goods [4]int

func (g goods) GetWidth() int {
	return g[1]
}

func (g goods) GetHeight() int {
	return g[2]
}

func (g goods) GetDepth() int {
	return g[3]
}

func (g goods) GetWeight() int {
	return 10
}

func TestPack(t *testing.T) {
	allBoxes := []Box[string]{
		{Width: 220, Height: 160, Depth: 100, Weight: 110, Item: "Box1"},
		{Width: 260, Height: 145, Depth: 145, Weight: 120, Item: "Box2"},
		{Width: 270, Height: 185, Depth: 110, Weight: 140, Item: "Box3"},
		{Width: 310, Height: 220, Depth: 140, Weight: 210, Item: "Box4"},
		{Width: 300, Height: 210, Depth: 200, Weight: 250, Item: "Box5"},
		{Width: 300, Height: 300, Depth: 130, Weight: 290, Item: "Box6"},
		{Width: 370, Height: 270, Depth: 150, Weight: 300, Item: "Box7"},
		{Width: 300, Height: 300, Depth: 250, Weight: 360, Item: "Box8"},
		{Width: 470, Height: 280, Depth: 210, Weight: 400, Item: "Box9"},
		{Width: 430, Height: 315, Depth: 200, Weight: 430, Item: "Box10"},
		{Width: 330, Height: 330, Depth: 350, Weight: 500, Item: "Box11"},
		{Width: 465, Height: 350, Depth: 370, Weight: 650, Item: "Box12"},
	}
	items := []Item{
		goods{1, 20, 100, 30},
		goods{2, 100, 20, 30},
		goods{3, 20, 100, 30},
		goods{4, 100, 20, 30},
		goods{5, 100, 20, 30},
		goods{6, 100, 100, 30},
		goods{7, 100, 100, 30},
	}
	want := []Box[string]{allBoxes[0]}
	want[0].Items = []BoxItem{
		{Item: items[5], RType: 0, Pos: [3]int{0, 0, 0}},
		{Item: items[6], RType: 0, Pos: [3]int{100, 0, 0}},
		{Item: items[0], RType: 0, Pos: [3]int{200, 0, 0}},
		{Item: items[1], RType: 0, Pos: [3]int{0, 100, 0}},
		{Item: items[2], RType: 1, Pos: [3]int{100, 100, 0}},
		{Item: items[3], RType: 2, Pos: [3]int{200, 100, 0}},
		{Item: items[4], RType: 0, Pos: [3]int{0, 120, 0}},
	}

	got, err := Pack(allBoxes, items)
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got:\n%swant:\n%s", printBoxes(got), printBoxes(want))
	}
}

func printBoxes[T any](boxes []Box[T]) (r string) {
	for i, box := range boxes {
		r += fmt.Sprintln("box", i, box.Width, box.Height, box.Depth, len(box.Items))
		for i, item := range box.Items {
			r += fmt.Sprintln("  ", i, item)
		}
	}

	return
}
