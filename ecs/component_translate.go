package ecs

type TranslateComponent struct {
	parent *Entity
	X, Y   float64
}

func NewTranslateComponent(parent *Entity) *TranslateComponent {
	return parent.addComponent(&TranslateComponent{
		parent: parent,
		X:      0,
		Y:      0,
	}).(*TranslateComponent)
}

func (t *TranslateComponent) Update(_ float64) {
	// empty
}

func (t *TranslateComponent) Type() ComponentType {
	return TranslateComponentType
}

func (t *TranslateComponent) Parent() *Entity {
	return t.parent
}
