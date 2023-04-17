package instance

import "lingmu/game-server/business/module/condition"

type Level struct {
	condition.Base
	Data int32
}
