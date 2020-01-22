package framework

type Manager struct {
	showCase map[string]Product
}

func NewManager() *Manager {
	return &Manager{
		showCase: map[string]Product{},
	}
}

func (m *Manager) Register(name string, proto Product) {
	m.showCase[name] = proto
}

func (m *Manager) Create(protoName string) Product {
	return m.showCase[protoName].CreateClone()
}
