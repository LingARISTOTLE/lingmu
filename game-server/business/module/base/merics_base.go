package base

/*
MetricsBase
@Description: 指标基础类
*/
type MetricsBase struct {
	Name string
}

func (m *MetricsBase) GetName() string {
	return m.Name
}

func (m *MetricsBase) SetName(str string) {
	m.Name = str
}

func (m *MetricsBase) Description() string {
	//TODO implement me
	panic("implement me")
}
