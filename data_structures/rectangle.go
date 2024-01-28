package data_structures

const (
	TileSize     = 16
	TextureWidth = 256
)

type Rect struct {
	X, Y, W, H int
}

func NewRectFromTileId(tileId int) *Rect {
	return &Rect{
		X: (tileId * TileSize) % TextureWidth,
		Y: (tileId * TileSize) / TextureWidth * TileSize,
		W: TileSize,
		H: TileSize,
	}
}
