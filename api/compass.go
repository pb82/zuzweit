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
	rad := math.Mod(c.radians, 2*math.Pi)
	if rad < 0 {
		rad += 2 * math.Pi
	}

	deg := rad * 180 / math.Pi

	switch {
	case deg >= 45 && deg < 135:
		return East
	case deg >= 135 && deg < 225:
		return South
	case deg >= 225 && deg < 315:
		return West
	default:
		return North
	}
}

func (c *Compass) TurnLeft() float64 {
	c.radians -= math.Pi / 2
	return c.radians
}

func (c *Compass) TurnRight() float64 {
	c.radians += math.Pi / 2
	return c.radians
}
