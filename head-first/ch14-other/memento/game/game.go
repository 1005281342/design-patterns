package game

type MasterGameObject struct {
	gameState string
}

func NewMasterGameObject() *MasterGameObject {
	return &MasterGameObject{gameState: "1"}
}

func (m *MasterGameObject) SetState(state string) {
	m.gameState = state
}

func (m *MasterGameObject) GetCurrentState() *Memento {
	return &Memento{state: m.gameState}
}

func (m *MasterGameObject) RestoreState(savedState *Memento) {
	m.gameState = savedState.SavedGameState()
}
