package util

import (
	"github.com/pb82/mini3d/api"
	"log"
)

// CubeWithTexture returns a cube with the same texture on each side.
func CubeWithTexture(x, y float64, atlas api.TextureAtlas) *api.Mesh {
	mesh := api.NewMesh()

	w := float64(atlas.W())
	h := float64(atlas.H())

	stepX := w / 2
	stepY := h / 2

	minX := x * stepX
	maxX := minX + stepX

	minY := y * stepY
	maxY := minY + stepY

	normalizedMinX := minX / w
	normalizedMaxX := maxX / w
	normalizedMinY := minY / h
	normalizedMaxY := maxY / h

	log.Println(normalizedMinX, normalizedMaxX, normalizedMinY, normalizedMaxY)

	triangle1 := api.Matrix3x2{
		{normalizedMinX, normalizedMaxY},
		{normalizedMinX, normalizedMinY},
		{normalizedMaxX, normalizedMinY},
	}

	triangle2 := api.Matrix3x2{
		{normalizedMinX, normalizedMaxY},
		{normalizedMaxX, normalizedMinY},
		{normalizedMaxX, normalizedMaxY},
	}

	// Side 1
	mesh.AddTriangle(api.TexturedTriangleFromMatrix(
		api.Matrix3x3{
			{0, 0, 0},
			{0, 1, 0},
			{1, 1, 0},
		},
		triangle1))

	mesh.AddTriangle(api.TexturedTriangleFromMatrix(
		api.Matrix3x3{
			{0, 0, 0},
			{1, 1, 0},
			{1, 0, 0},
		},
		triangle2))

	// Side 2
	mesh.AddTriangle(api.TexturedTriangleFromMatrix(
		api.Matrix3x3{
			{1, 0, 0},
			{1, 1, 0},
			{1, 1, 1},
		},
		triangle1))

	mesh.AddTriangle(api.TexturedTriangleFromMatrix(
		api.Matrix3x3{
			{1, 0, 0},
			{1, 1, 1},
			{1, 0, 1},
		},
		triangle2))

	// Side 3
	mesh.AddTriangle(api.TexturedTriangleFromMatrix(
		api.Matrix3x3{
			{1, 0, 1},
			{1, 1, 1},
			{0, 1, 1},
		},
		triangle1))

	mesh.AddTriangle(api.TexturedTriangleFromMatrix(
		api.Matrix3x3{
			{1, 0, 1},
			{0, 1, 1},
			{0, 0, 1},
		},
		triangle2))

	// Side 4
	mesh.AddTriangle(api.TexturedTriangleFromMatrix(
		api.Matrix3x3{
			{0, 0, 1},
			{0, 1, 1},
			{0, 1, 0},
		},
		triangle1))

	mesh.AddTriangle(api.TexturedTriangleFromMatrix(
		api.Matrix3x3{
			{0, 0, 1},
			{0, 1, 0},
			{0, 0, 0},
		},
		triangle2))

	// Side 5
	mesh.AddTriangle(api.TexturedTriangleFromMatrix(
		api.Matrix3x3{
			{0, 1, 0},
			{0, 1, 1},
			{1, 1, 1},
		},
		triangle1))

	mesh.AddTriangle(api.TexturedTriangleFromMatrix(
		api.Matrix3x3{
			{0, 1, 0},
			{1, 1, 1},
			{1, 1, 0},
		},
		triangle2))

	// Side 6
	mesh.AddTriangle(api.TexturedTriangleFromMatrix(
		api.Matrix3x3{
			{1, 0, 1},
			{0, 0, 1},
			{0, 0, 0},
		},
		triangle1))

	mesh.AddTriangle(api.TexturedTriangleFromMatrix(
		api.Matrix3x3{
			{1, 0, 1},
			{0, 0, 0},
			{1, 0, 0},
		},
		triangle2))

	return mesh
}
