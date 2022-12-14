package testapi

import (
	"context"
	"sync"

	"github.com/vanti-dev/sc-bos/pkg/gen"
)

type API struct {
	gen.UnimplementedTestApiServer

	m    sync.RWMutex
	data string
}

func (api *API) GetTest(_ context.Context, _ *gen.GetTestRequest) (*gen.Test, error) {
	api.m.RLock()
	defer api.m.RUnlock()

	data := api.data
	return &gen.Test{Data: data}, nil
}

func (api *API) UpdateTest(_ context.Context, request *gen.UpdateTestRequest) (*gen.Test, error) {
	api.m.Lock()
	defer api.m.Unlock()

	data := request.GetTest().GetData()
	api.data = data
	return &gen.Test{Data: data}, nil
}

func NewAPI() *API {
	api := &API{
		data: "default data",
	}

	return api
}