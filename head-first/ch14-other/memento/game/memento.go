package game

type Memento struct {
	state string
}

// SavedGameState 已保存的游戏状态
func (m *Memento) SavedGameState() string {
	return m.state
}
