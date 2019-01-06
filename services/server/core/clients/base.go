package clients

import (
	"context"
	"sync"

	api "github.com/ilovelili/dongfeng-physique/services/proto"
	"github.com/ilovelili/dongfeng-physique/services/utils"
	"github.com/micro/go-micro/client"
)

var instance *Client
var once sync.Once

// Client struct represents InvastBroker API client
type Client struct {
	ServiceName string
	client      api.ApiService
}

// New creates a new API struct
func New() *Client {
	once.Do(func() {
		config := utils.GetConfig()
		servicename := config.ServiceNames.CoreServer
		instance = &Client{
			ServiceName: servicename,
			client:      api.NewApiService(servicename, client.DefaultClient), // rpc client calling core server
		}
	})

	return instance
}

func ctx() context.Context {
	return context.Background()
}
