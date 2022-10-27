package merge

import (
	"github.com/smart-core-os/sc-golang/pkg/trait"
	"github.com/vanti-dev/bsp-ew/internal/driver/bacnet/config"
	"github.com/vanti-dev/bsp-ew/internal/driver/bacnet/known"
	"github.com/vanti-dev/bsp-ew/internal/node"
	"github.com/vanti-dev/gobacnet"
)

func IntoTrait(client *gobacnet.Client, ctx known.Context, traitConfig config.RawTrait) (node.SelfAnnouncer, error) {
	// todo: implement some traits that pull data from different bacnet devices.
	switch traitConfig.Kind {
	case trait.FanSpeed:
	case trait.AirTemperature:
	}
	return nil, ErrTraitNotSupported
}