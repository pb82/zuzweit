package ecs

type AdvanceForward struct {
	player *Entity
}

func (a *AdvanceForward) Run(dt float64) {
	translate := a.player.GetComponent(TranslateComponentType).(*TranslateComponent)
	translate.Y = translate.Y + 1*(dt*1000)
}

func (a *AdvanceForward) Complete() bool {
	return false
}

func NewAdvance(player *Entity) *AdvanceForward {
	return &AdvanceForward{player}
}
