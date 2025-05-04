package api

import (
	"math"
)

type Direction int

const (
	North Direction = iota
	East
	South
	West
	Invalid
)

func (d Direction) String() string {
	switch d {
	case North:
		return "North"
	case East:
		return "East"
	case South:
		return "South"
	case West:
		return "West"
	case Invalid:
		return "Invalid"
	default:
		return "Invalid"
	}
}

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

func (c *Compass) GetDirection() Direction {
	steps := int(math.Abs(c.radians) / (math.Pi / 2))
	options := []Direction{North, East, South, West}
	return options[steps%len(options)]
}

func (c *Compass) TurnLeft() float64 {
	c.radians -= math.Pi / 2
	return c.radians
}

func (c *Compass) TurnRight() float64 {
	c.radians += math.Pi / 2
	return c.radians
}
