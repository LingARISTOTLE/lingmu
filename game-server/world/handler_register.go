package world

func (m *ManagerHost) HandlerRegister() {
	m.Handlers[1] = m.UserLogin
}
