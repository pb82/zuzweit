package api

import "math"

type Map struct {
	data [][]int
	W, H int
}

func NewMap(w, h int) *Map {
	data := make([][]int, h)
	for i := range data {
		data[i] = make([]int, w)
	}

	return &Map{
		W:    w,
		H:    h,
		data: data,
	}
}

func NewDemoMap() *Map {
	return &Map{
		data: [][]int{
			{1, 1, 1, 1, 1, 1},
			{1, 0, 1, 0, 0, 1},
			{1, 0, 1, 0, 0, 1},
			{1, 0, 1, 1, 0, 1},
			{1, 0, 0, 0, 2, 1},
			{1, 1, 1, 1, 1, 1}},
		W: 6,
		H: 6,
	}
}

func (m *Map) Get(x, y float64) int {
	_x := int(math.Floor(x))
	_y := int(math.Floor(y))

	if len(m.data) <= _y || len(m.data[_y]) <= _x {
		return -1
	}

	return m.data[_y][_x]
}

func (m *Map) IsBlocked(x, y float64) bool {
	return m.Get(x, y) == 1
}

func (m *Map) GetPlayerStart() (float64, float64) {
	for y := 0; y < m.H; y++ {
		for x := 0; x < m.W; x++ {
			if m.data[y][x] == 2 {
				return float64(x) + .5, float64(y) + .5
			}
		}
	}

	return 0, 0
}
