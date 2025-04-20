package api

import (
	"zuzweit/data_structures"
)

type Command interface {
	Complete() bool
	Run(dt float64)
}

var instance *data_structures.Queue[Command]

func GetCommandQueue() *data_structures.Queue[Command] {
	if instance == nil {
		instance = new(data_structures.Queue[Command])
	}
	return instance
}
