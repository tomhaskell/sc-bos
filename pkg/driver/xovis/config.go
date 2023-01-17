package xovis

import (
	"github.com/vanti-dev/sc-bos/pkg/driver"
)

type DriverConfig struct {
	driver.BaseConfig
	MultiSensor bool            `json:"multiSensor"`
	Host        string          `json:"host"`
	Username    string          `json:"username"`
	Password    string          `json:"password"`
	DataPush    *DataPushConfig `json:"dataPush"`
	Devices     []DeviceConfig  `json:"devices,omitempty"`
}

type DeviceConfig struct {
	Name       string       `json:"name"`
	Occupancy  *LogicConfig `json:"occupancy"`  // an Occupancy logic
	EnterLeave *LogicConfig `json:"enterLeave"` // an In/Out logic
}

type LogicConfig struct {
	ID int `json:"id"`
}

type DataPushConfig struct {
	WebhookPath string `json:"webhookPath"`
}
