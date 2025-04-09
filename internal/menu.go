package internal

import "zuzweit/data_structures"

type MenuCallback func(m *MenuRoot)

type MenuItem struct {
	Title   string
	OnClick MenuCallback
}

type Menu struct {
	Items []*MenuItem
}

type MenuRoot struct {
	menus data_structures.Stack[*Menu]
}

func (m *MenuRoot) PushMenu(menu *Menu) {
	m.menus.Push(menu)
}

func (m *MenuRoot) Pop() {
	m.menus.Pop()
}

func (m *MenuRoot) Top() (*Menu, error) {
	return m.menus.Peek()
}
