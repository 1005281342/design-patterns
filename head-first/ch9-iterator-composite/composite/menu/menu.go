package menu

import "log"

type Menu struct {
	*Item
	subMenus []Component
}

func NewMenu(name string, desc string) *Menu {
	return &Menu{Item: &Item{name: name, desc: desc}}
}

func (m *Menu) Print() {
	log.Printf("Menu name:%s desc:%s\n", m.GetName(), m.GetDesc())
	for i, subMenu := range m.subMenus {
		log.Printf("%s subMenu[%d]: ", m.GetName(), i)
		subMenu.Print()
	}
}

func (m *Menu) Add(c Component) {
	m.subMenus = append(m.subMenus, c)
}

func (m *Menu) Remove(c Component) {
	for i, subMenu := range m.subMenus {
		if subMenu != c {
			continue
		}

		m.subMenus = append(m.subMenus[:i], m.subMenus[i+1:]...)
		return
	}
}

func (m *Menu) GetChild(index int) Component {
	if index >= len(m.subMenus) {
		return nil
	}

	return m.subMenus[index]
}
