package logger

import (
	"github.com/LingARISTOTLE/lingolog"
	"log"
	"sync"
)

var (
	Logger         *spoor.Spoor
	onceInitLogger sync.Once
)

/*
init
@Description: 初始化日志打印配置
*/
func init() {
	onceInitLogger.Do(func() {
		fileWriter := spoor.NewFileWriter("log", 0, 0, 0)
		newSpoor := spoor.NewSpoor(spoor.DEBUG, "", log.Ldate|log.Ltime|log.Lmicroseconds|log.Llongfile, spoor.WithFileWriter(fileWriter))
		Logger = newSpoor
	})
}
