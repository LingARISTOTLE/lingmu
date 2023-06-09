package gamedemo

func main() {
	LoadConf()
	LoadDBData()
	StartGameServer()
	ResolveGameLogic()
	StopGameServer()
}

/*
StartGameServer 启动游戏服务
*/
func StartGameServer() {
	Log()
}

/*
LoadConf 加载配置
*/
func LoadConf() {
	Log()
}

/*
LoadDBData  加载DB数据
*/
func LoadDBData() {
	Log()
}

/*
ResolveGameLogic 处理游戏逻辑
*/
func ResolveGameLogic() {
	SaveGameData()
	Log()
}

/*
SaveGameData 存储游戏数据
*/
func SaveGameData() {
	Log()
}

/*
StopGameServer 停止游戏服务器
*/
func StopGameServer() {
	Log()
}

/*
Log 打印日志
*/
func Log() {

}
