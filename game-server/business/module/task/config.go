package task

type Config struct {
	Id       uint32        `json:"id"`
	Name     string        `json:"name"`
	DropId   uint32        `json:"dropId"`
	Category int           `json:"category"` //类别
	Targets  []*TargetConf `json:"targets"`
}

type TargetConf struct {
}
