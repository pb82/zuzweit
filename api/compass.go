package api

import "math"

type Compass struct {
	radians float64
}

func NewCompass() *Compass {
	return &Compass{
		radians: 0,
	}
}

func (c *Compass) GetRadians() float64 {
	return c.radians
}

func (c *Compass) TurnLeft() float64 {
	c.radians -= math.Pi / 2
	return c.radians
}

func (c *Compass) TurnRight() float64 {
	c.radians += math.Pi / 2
	return c.radians
}
