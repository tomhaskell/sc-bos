package trv

import (
	"context"

	"go.uber.org/zap"

	"github.com/smart-core-os/sc-golang/pkg/trait"
	"github.com/smart-core-os/sc-golang/pkg/trait/airtemperature"
	"github.com/vanti-dev/sc-bos-drivers/pkg/driver/shelly/trv/config"
	"github.com/vanti-dev/sc-bos/pkg/driver"
	"github.com/vanti-dev/sc-bos/pkg/node"
	"github.com/vanti-dev/sc-bos/pkg/task/service"
)

const DriverName = "vanti/shelly/trv"

var Factory driver.Factory = factory{}

type factory struct{}

func (f factory) New(services driver.Services) service.Lifecycle {
	d := &Driver{
		announcer: services.Node,
	}
	d.Service = service.New(service.MonoApply(d.applyConfig))
	d.logger = services.Logger.Named(DriverName)
	return d
}

type Driver struct {
	*service.Service[config.Root]
	logger    *zap.Logger
	announcer node.Announcer

	devices []*TRV
}

func (d *Driver) applyConfig(ctx context.Context, cfg config.Root) error {
	announcer := node.AnnounceContext(ctx, d.announcer)

	for _, device := range cfg.Devices {
		logger := d.logger.Named(device.Name)
		trv, err := NewTRV(device, logger)
		if err != nil {
			d.logger.Warn("Problem adding device: " + err.Error())
		}
		d.devices = append(d.devices, trv)

		announcer.Announce(device.Name, node.HasTrait(trait.AirTemperature, node.WithClients(airtemperature.WrapApi(trv.airTemperatureServer))))
	}

	return nil
}
