// Package airthings integrates AirThings devices into Smart Core.
// The primary api used by this driver is the location latest samples api.
//
// The driver pulls all data into a local model, then translates that local model into Smart Core traits.
package airthings

import (
	"context"
	"net/http"
	"sync"

	"golang.org/x/sync/errgroup"

	"github.com/vanti-dev/sc-bos/pkg/driver"
	"github.com/vanti-dev/sc-bos/pkg/driver/airthings/api"
	"github.com/vanti-dev/sc-bos/pkg/driver/airthings/local"
	"github.com/vanti-dev/sc-bos/pkg/gen"
	"github.com/vanti-dev/sc-bos/pkg/gentrait/statuspb"
	"github.com/vanti-dev/sc-bos/pkg/node"
	"github.com/vanti-dev/sc-bos/pkg/task/service"
)

const DriverName = "airthings"

var Factory driver.Factory = factory{}

type factory struct{}

func (f factory) New(services driver.Services) service.Lifecycle {
	services.Logger = services.Logger.Named(DriverName)
	d := &Driver{
		Services: services,
	}
	d.Service = service.New(service.MonoApply(d.applyConfig))
	return d
}

type Driver struct {
	*service.Service[Config]
	driver.Services

	cfg    Config
	client *http.Client

	listLocationsOnce sync.Once
	locationsErr      error
	locations         api.GetLocationsResponse
}

func (d *Driver) applyConfig(ctx context.Context, cfg Config) error {
	announcer := node.AnnounceContext(ctx, d.Node)
	d.listLocationsOnce = sync.Once{}
	d.cfg = cfg

	ccConfig, err := d.cfg.Auth.ClientCredentialsConfig()
	if err != nil {
		return err
	}
	d.client = ccConfig.Client(ctx)

	status := statuspb.NewMap(announcer)

	grp, ctx := errgroup.WithContext(ctx)
	for _, location := range cfg.Locations {
		location := location
		ll := local.NewLocation()
		grp.Go(func() error {
			return d.pollLocationLatestSamples(ctx, location, ll)
		})

		for _, device := range location.Devices {
			n := device.Name
			announcer.Announce(n, node.HasMetadata(device.Metadata))
			status.UpdateProblem(n, &gen.StatusLog_Problem{
				Level:       gen.StatusLog_NOMINAL,
				Description: "Device configured successfully",
				Name:        n + ":setup",
			})
			err = d.announceDevice(ctx, announcer, device, ll, status)
			if err != nil {
				return err // failure of configuration, not runtime
			}
		}
	}
	return nil
}
