package api

import "math"

type CardinalDirection int

const (
	North CardinalDirection = iota
	South
	East
	West
)

func (c CardinalDirection) TurnLeft() CardinalDirection {
	switch c {
	case North:
		return West
	case South:
		return East
	case East:
		return North
	case West:
		return South
	}

	return North
}

func (c CardinalDirection) TurnRight() CardinalDirection {
	switch c {
	case North:
		return East
	case South:
		return West
	case East:
		return South
	case West:
		return North
	}

	return North
}

type Compass struct {
	radians   float64
	direction CardinalDirection
}

func NewCompass() *Compass {
	return &Compass{
		radians:   0,
		direction: East,
	}
}

func (c *Compass) GetCardinalDirection() CardinalDirection {
	return c.direction
}

func (c *Compass) GetRadians() float64 {
	return c.radians
}

func (c *Compass) SetCardinalDirection(direction CardinalDirection) {
	switch direction {
	case North:
		c.direction = North
		c.radians = math.Pi / 2
	case South:
		c.direction = South
		c.radians = (3 * math.Pi) / 2
	case East:
		c.direction = East
		c.radians = 0
	case West:
		c.direction = West
		c.radians = math.Pi
	}
}

func (c *Compass) TurnLeft() float64 {
	c.SetCardinalDirection(c.direction.TurnLeft())
	return c.radians
}

func (c *Compass) TurnRight() float64 {
	c.SetCardinalDirection(c.direction.TurnRight())
	return c.radians
}
