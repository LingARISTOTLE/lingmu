package mongo

import (
	"context"
	"github.com/LingARISTOTLE/broker"
	mongobrocker "github.com/LingARISTOTLE/broker/mongo"
	"sync"
)

// mongodb连接工具包
var (
	Client        *mongobrocker.Client
	onceInitMongo sync.Once
)

func init() {
	onceInitMongo.Do(func() {
		ctx := context.Background()
		tc := &mongobrocker.Client{
			BaseComponent: broker.NewBaseComponent(),
			RealCli: mongobrocker.NewClient(ctx, &mongobrocker.Config{
				URI:         "mongodb://47.120.7.164:27017",
				MinPoolSize: 3,
				MaxPoolSize: 3000,
			}),
		}

		tc.Launch()
		defer tc.Stop()
	})
}
