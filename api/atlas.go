package api

import (
	"bytes"
	"image"
	"image/color"

	_ "image/png"
)

type TextureAtlasImpl struct {
	w, h int
	img  image.Image
}

func (t *TextureAtlasImpl) W() int {
	return t.w
}

func (t *TextureAtlasImpl) H() int {
	return t.h
}

func (t *TextureAtlasImpl) ColorAt(x, y int) color.Color {
	return t.img.At(x, y)
}

func (t *TextureAtlasImpl) LoadTexture(textureData []byte) {
	texture, _, _ := image.Decode(bytes.NewReader(textureData))
	t.img = texture
	t.w = texture.Bounds().Dx()
	t.h = texture.Bounds().Dy()
}
