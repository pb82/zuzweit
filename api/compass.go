package api

import (
	"log"
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
	increment := math.Pi / 4
	steps := int(math.Abs(c.radians) / increment)
	positions := [4]Direction{North, East, South, West}

	log.Println(c.radians, steps)

	return positions[steps%4]
}

func (c *Compass) TurnLeft() float64 {
	c.radians -= math.Pi / 2
	return c.radians
}

func (c *Compass) TurnRight() float64 {
	c.radians += math.Pi / 2
	return c.radians
}
