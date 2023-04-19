package launch

import (
	"lingmu/game-server/aop/pprof"
	"lingmu/game-server/aop/pprof/web"
)

func main() {
	h := pprof.Handler{
		Router: web.NewHttpRouter(),
	}
	h.RegisterProfiler()

}
