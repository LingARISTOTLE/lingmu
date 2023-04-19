package launch

import (
	"lingmu/game-server/business/server/gateway/client"
	"lingmu/game-server/business/server/gateway/server"
	"time"
)

/*
Gateway
@Description: 网关本体
*/
type Gateway struct {
	clientServer client.Server //客户端
	serverServer server.Server //服务器
	Metrics      *MetricInfo   //指标
	Config       Conf          //配置
}

/*
MetricInfo
@Description: 指标实体
*/
type MetricInfo struct {
	ClientCount int32 `json:"client_count"`
	ServerCount int32 `json:"server_count"`
}

/*
Update
@Description: 将服务和客户的时时数量同步到指标中
@receiver g
*/
func (g *Gateway) Update() {
	g.Metrics.ServerCount = g.serverServer.GetServerCount()
	g.Metrics.ClientCount = g.clientServer.GetServerCount()
}

func (g *Gateway) Loop() {

	tick := time.NewTicker(time.Duration(g.Config.UpdateInfoCd) * time.Second)

	for {
		select {
		case <-tick.C:
			g.Update()
		}
	}
}
